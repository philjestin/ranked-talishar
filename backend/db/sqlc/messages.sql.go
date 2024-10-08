// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: messages.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const getConversationsByUser = `-- name: GetConversationsByUser :many
select c.id, c.last_message, c.last_message_timestamp, c.last_message_sender_id, c.message_count, c.created_at
FROM conversations c
JOIN participants p ON c.id = p.conversation_id
WHERE p.user_id = $1
ORDER BY c.created_at DESC
`

func (q *Queries) GetConversationsByUser(ctx context.Context, userID uuid.NullUUID) ([]Conversation, error) {
	rows, err := q.query(ctx, q.getConversationsByUserStmt, getConversationsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Conversation{}
	for rows.Next() {
		var i Conversation
		if err := rows.Scan(
			&i.ID,
			&i.LastMessage,
			&i.LastMessageTimestamp,
			&i.LastMessageSenderID,
			&i.MessageCount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMessagesByConversation = `-- name: GetMessagesByConversation :many
SELECT id, conversation_id, sender_id, sender_username, content, created_at
FROM messages
where conversation_id = $1
ORDER BY created_at ASC
`

func (q *Queries) GetMessagesByConversation(ctx context.Context, conversationID sql.NullInt32) ([]Message, error) {
	rows, err := q.query(ctx, q.getMessagesByConversationStmt, getMessagesByConversation, conversationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Message{}
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.ConversationID,
			&i.SenderID,
			&i.SenderUsername,
			&i.Content,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const sendMessage = `-- name: SendMessage :one
INSERT INTO messages (conversation_id, sender_id, sender_username, content, created_at)
VALUES ($1, $2, $3, $4, NOW())
RETURNING id, conversation_id, sender_id, sender_username, content, created_at
`

type SendMessageParams struct {
	ConversationID sql.NullInt32  `json:"conversation_id"`
	SenderID       uuid.NullUUID  `json:"sender_id"`
	SenderUsername sql.NullString `json:"sender_username"`
	Content        string         `json:"content"`
}

func (q *Queries) SendMessage(ctx context.Context, arg SendMessageParams) (Message, error) {
	row := q.queryRow(ctx, q.sendMessageStmt, sendMessage,
		arg.ConversationID,
		arg.SenderID,
		arg.SenderUsername,
		arg.Content,
	)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.ConversationID,
		&i.SenderID,
		&i.SenderUsername,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}
