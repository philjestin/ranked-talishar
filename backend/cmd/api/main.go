package main

import (
	"context"
	"database/sql"
	"expvar"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq"
	dbCon "github.com/philjestin/ranked-talishar/db/sqlc"
	"github.com/philjestin/ranked-talishar/internal/data"
	"github.com/philjestin/ranked-talishar/internal/mailer"
	"github.com/philjestin/ranked-talishar/internal/vcs"
	"github.com/philjestin/ranked-talishar/listener"
)

type appConfig struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
	limiter struct {
		enabled bool
		rps     float64
		burst   int
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}

	cors struct {
		trustedOrigins []string
	}
}

type application struct {
	config appConfig
	logger *slog.Logger
	wg     sync.WaitGroup
	models data.Models
	mailer mailer.Mailer
}

var (
	version = vcs.Version()
)

func main() {
	var cfg appConfig

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.db.dsn, "db-dsn", "", "PostgreSQL DSN")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgreSQL max connection idle time")

	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")
	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 25, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "a7420fc0883489", "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "e75ffd0a3aa5ec", "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "Greenlight <no-reply@github.com/philjestin/ranked-talishar-pro>", "SMTP sender")

	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		cfg.cors.trustedOrigins = strings.Fields(val)
		return nil
	})

	displayVersion := flag.Bool("version", false, "Display version and exit")

	flag.Parse()

	if *displayVersion {
		fmt.Printf("Version:\t%s\n", version)
		os.Exit(0)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, conn, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// Create context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cancellation on exit

	// Start notification listener in a separate goroutine
	go func() {
		listenerErr := listener.ListenNotifications(ctx, cfg.db.dsn, "update_ratings_channel", db)
		if listenerErr != nil {
			log.Fatal(listenerErr)
		}
	}()

	logger.Info("database connection pool established")

	expvar.NewString("version").Set(version)

	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	expvar.Publish("database", expvar.Func(func() any {
		return conn.Stats()
	}))

	expvar.Publish("timestamp", expvar.Func(func() any {
		return time.Now().Unix()
	}))

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	err = app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(cfg appConfig) (*dbCon.Queries, *sql.DB, error) {
	// Open the database connection
	conn, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	conn.SetMaxOpenConns(cfg.db.maxOpenConns)
	conn.SetMaxIdleConns(cfg.db.maxIdleConns)
	conn.SetConnMaxIdleTime(cfg.db.maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = conn.PingContext(ctx)
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	// Initialize the Queries object
	db := dbCon.New(conn)
	return db, conn, nil
}
