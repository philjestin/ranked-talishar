ALTER TABLE users
ADD COLUMN IF NOT EXISTS activated bool NOT NULL default false,
ADD COLUMN IF NOT EXISTS version integer NOT NULL DEFAULT 1;