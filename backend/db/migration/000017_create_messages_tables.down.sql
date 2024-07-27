DROP TABLE IF EXISTS conversations;
DROP TABLE IF EXISTS participants;
DROP TABLE IF EXISTS messages;


DROP TRIGGER IF EXISTS trg_update_conversation_last_message;
DROP FUNCTION IF EXISTS update_conversation_last_message();