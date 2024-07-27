-- name: CreateConversation :one
INSERT INTO conversations (created_at)
VALUES(NOW())
RETURNING id, created_at;

-- name: AddParticipant :exec
INSERT INTO participants (conversation_id, user_id, joined_at)
VALUES ($1, $2, NOW());
