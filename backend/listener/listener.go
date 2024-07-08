package listener

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	db "github.com/philjestin/ranked-talishar/db/sqlc"
	"github.com/philjestin/ranked-talishar/elo"

	_ "github.com/lib/pq" // Assuming you're using PostgreSQL
)

// NotificationPayload represents the data structure of your notification message
type NotificationPayload struct {
	MatchID uuid.UUID `json:"match_id"`
}

func ListenNotifications(ctx context.Context, connString string, channel string, q *db.Queries) error {
	log.Println("Attempting to start the notification listener...")

	listener := pq.NewListener(connString, 10*time.Second, time.Minute, nil)
	err := listener.Listen(channel)
	if err != nil {
		log.Printf("Listener error: %v\n", err)
		return err
	}
	defer listener.Close()

	log.Printf("Listening on channel: %s\n", channel) // Log to confirm listening on the channel

	for {
		select {
		case <-ctx.Done():
			log.Println("context canceled, stopping listener")
			return nil
		case notification := <-listener.Notify:
			if notification == nil {
				continue
			}

			// Handle notification payload
			var payload NotificationPayload
			if err := json.Unmarshal([]byte(notification.Extra), &payload); err != nil {
				log.Printf("Error unmarshalling notification payload: %v\n", err)
				continue
			}

			// Trigger ELO update based on match ID
			go handleMatchCompletion(payload.MatchID, q)
		}
	}
}

// handleMatchCompletion retrieves player information and updates ratings based on match ID
func handleMatchCompletion(matchID uuid.UUID, q *db.Queries) {
	// Update player ratings based on match ID
	err := updateMatchRatingsFromID(q, matchID)
	if err != nil {
		log.Printf("Error updating ratings for match %d: %v\n", matchID, err)
		return
	}

	log.Printf("Successfully updated ratings for match %d\n", matchID)
}

// updateMatchRatingsFromID retrieves winner and loser IDs and updates their ratings
func updateMatchRatingsFromID(q *db.Queries, matchID uuid.UUID) error {
	// Retrieve winner and loser IDs based on match ID
	winnerID, loserID, err := getPlayersFromMatch(q, matchID)
	if err != nil {
		return err
	}

	// Update player ratings
	err = elo.UpdateRatings(context.Background(), q, winnerID, loserID, 1.0)
	if err != nil {
		return err
	}

	return nil
}

// getPlayersFromMatch retrieves winner and loser IDs based on match ID
func getPlayersFromMatch(q *db.Queries, matchID uuid.UUID) (uuid.UUID, uuid.UUID, error) {
	matchPlayers, err := q.GetMatchPlayers(context.Background(), matchID)
	if err != nil {
		log.Fatalf("Error getting match players: %v", err)
	}

	var winnerID, loserID uuid.UUID
	if len(matchPlayers) > 0 {
		winnerID = matchPlayers[0].WinnerID.UUID
		loserID = matchPlayers[0].LoserID.UUID
	} else {
		log.Println("No match players found")
	}

	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	return winnerID, loserID, nil
}
