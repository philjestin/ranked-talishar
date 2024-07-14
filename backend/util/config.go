package util

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	db "github.com/philjestin/ranked-talishar/db/sqlc"
	"github.com/spf13/viper"
)

type Config struct {
	DbDriver            string        `mapstructure:"DB_DRIVER"`
	DbSource            string        `mapstructure:"DB_SOURCE"`
	PostgresUser        string        `mapstructure:"POSTGRES_USER"`
	PostgresPassword    string        `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDb          string        `mapstructure:"POSTGRES_DB"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// SendMatchUpdateNotification sends a notification for a match update
func SendMatchUpdateNotification(ctx context.Context, db *db.Queries, matchID uuid.UUID, winnerId uuid.UUID, loserId uuid.UUID) error {
	log.Println("Sending Match Update Notification")

	// Load configuration
	config, err := LoadConfig("..")
	if err != nil {
		return fmt.Errorf("error loading configuration: %w", err)
	}

	dbUser := config.PostgresUser
	dbPassword := config.PostgresPassword
	dbName := config.PostgresDb

	payload := NotificationPayload{
		MatchID:  matchID,
		LoserID:  loserId,
		WinnerID: winnerId,
	}

	log.Println("The payload in the SendMatchUpdateNotification is %s", payload)

	connString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)

	return SendNotification(ctx, connString, "update_ratings_channel", payload)
}

// NotificationPayload represents the data structure of your notification message
type NotificationPayload struct {
	MatchID  uuid.UUID `json:"match_id"`
	WinnerID uuid.UUID `json:"winner_id"`
	LoserID  uuid.UUID `json:"loser_id"`
}

func SendNotification(ctx context.Context, connString string, channel string, payload interface{}) error {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return fmt.Errorf("failed to open connection: %w", err)
	}
	defer db.Close()

	// Convert payload to JSON bytes
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	log.Println("The payloadBytes in the SendNotification func is %s", string(payloadBytes))

	values := []interface{}{channel, payloadBytes}
	notificationString := fmt.Sprintf("NOTIFY %s, '%s'", values...)

	log.Println("The notification string is ", notificationString)

	// Send notification using pq.Notify
	_, err = db.ExecContext(ctx, notificationString)
	if err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}

	log.Printf("Sent notification on channel: %s with payload: %s\n", channel, string(payloadBytes))
	return nil
}
