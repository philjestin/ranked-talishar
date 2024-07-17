package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/philjestin/ranked-talishar/controllers"
	dbCon "github.com/philjestin/ranked-talishar/db/sqlc"
	"github.com/philjestin/ranked-talishar/listener"
	"github.com/philjestin/ranked-talishar/routes"
	"github.com/philjestin/ranked-talishar/token"
	"github.com/philjestin/ranked-talishar/util"
)

var (
	server *gin.Engine
	db     *dbCon.Queries

	ContactController controllers.ContactController
	ContactRoutes     routes.ContactRoutes
	UserController    controllers.UserController
	UserRoutes        routes.UserRoutes
	GameController    controllers.GameController
	GameRoutes        routes.GameRoutes
	FormatController  controllers.FormatController
	FormatRoutes      routes.FormatRoutes
	HeroController    controllers.HeroController
	HeroRoutes        routes.HeroRoutes
	MatchController   controllers.MatchController
	MatchRoutes       routes.MatchRoutes

	jwtMaker      token.Maker
	tokenDuration time.Duration
)

func init() {
	// Load the configuration
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	secretKey := config.TokenSymmetricKey
	if secretKey == "" {
		log.Fatal("Missing JWT secret key")
	}

	tokenDuration = config.AccessTokenDuration
	jwtMaker, err = token.NewJWTMaker(secretKey)
	if err != nil {
		log.Fatal("Failed to create JWT maker:", err)
	}

	// Open the database connection
	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Initialize the Queries object
	db = dbCon.New(conn)
	fmt.Println("PostgreSql connected successfully...", db)

	// Initialize controllers and routes
	ContactController = *controllers.NewContactController(db, context.Background())
	ContactRoutes = routes.NewRouteContact(ContactController)

	UserController = *controllers.NewUserController(db, context.Background(), jwtMaker, tokenDuration)
	UserRoutes = routes.NewRouteUser(UserController, jwtMaker)

	GameController = *controllers.NewGameController(db, context.Background())
	GameRoutes = routes.NewRouteGame(GameController)

	FormatController = *controllers.NewFormatController(db, context.Background())
	FormatRoutes = routes.NewRouteFormat(FormatController)

	HeroController = *controllers.NewHeroController(db, context.Background())
	HeroRoutes = routes.NewRouteHero(HeroController)

	MatchController = *controllers.NewMatchController(db, context.Background())
	MatchRoutes = routes.NewRouteMatch(MatchController)

	// Initialize the Gin server
	server = gin.Default()
}

func main() {
	// Use the configuration loaded in init
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Create context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cancellation on exit

	// Start notification listener in a separate goroutine
	go func() {
		listenerErr := listener.ListenNotifications(ctx, config.DbSource, "update_ratings_channel", db)
		if listenerErr != nil {
			log.Fatal(listenerErr)
		}
	}()

	// Set up the router group
	router := server.Group("/api")

	// Health check endpoint
	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "The ranked-talishar API is working fine"})
	})

	// Register routes
	ContactRoutes.ContactRoute(router)
	UserRoutes.UserRoute(router)
	GameRoutes.GameRoute(router)
	FormatRoutes.FormatRoute(router)
	HeroRoutes.HeroRoute(router)
	MatchRoutes.MatchRoute(router)

	// Handle 404 for undefined routes
	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": fmt.Sprintf("The specified route %s not found", ctx.Request.URL)})
	})

	// Start the server
	log.Fatal(server.Run(":" + config.ServerAddress))
}
