package notifier

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/philjestin/ranked-talishar/util"

	"github.com/lib/pq"
)

// NotificationPayload represents the data structure of your notification message
type NotificationPayload struct {
	MatchID uuid.UUID `json:"match_id"`
}

// SendNotification sends a notification to the specified channel with the provided payload
func SendNotification(ctx context.Context, config util.Config, channel string, payload NotificationPayload) error {
	if channel == "" {
		return errors.New("notification channel cannot be empty")
	}

	db, err := pq.Open(config.DbDriver, config.DbSource)
	if err != nil {
		return err
	}
	defer db.Close()

	// Marshal payload to JSON
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Build notification message
	message := map[string]interface{}{
		"type":  "update_ratings_channel",
		"extra": string(data),
	}

	// Send notification
	_, err = db.Notify(ctx, "NOTIFY ?, ?", message["type"], message["extra"])
	if err != nil {
		return err
	}

	return nil
}
