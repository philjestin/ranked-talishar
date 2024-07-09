ALTER TABLE users
ADD COLUMN IF NOT EXISTS wins INTEGER NOT NULL default 0,
ADD COLUMN IF NOT EXISTS losses INTEGER NOT NULL DEFAULT 0,
ADD COLUMN IF NOT EXISTS ties INTEGER NOT NULL DEFAULT 0,
ADD COLUMN IF NOT EXISTS elo INTEGER NOT NULL DEFAULT 1500;