-- name: SendMessage :one
INSERT INTO messages (conversation_id, sender_id, sender_username, content, created_at)
VALUES ($1, $2, $3, $4, NOW())
RETURNING id, conversation_id, sender_id, sender_username, content, created_at;

-- name: GetMessagesByConversation :many
SELECT id, conversation_id, sender_id, sender_username, content, created_at
FROM messages
where conversation_id = $1
ORDER BY created_at ASC;

-- name: GetConversationsByUser :many
select c.id, c.last_message, c.last_message_timestamp, c.last_message_sender_id, c.message_count, c.created_at
FROM conversations c
JOIN participants p ON c.id = p.conversation_id
WHERE p.user_id = $1
ORDER BY c.created_at DESC;