ALTER TABLE users
ADD COLUMN IF NOT EXISTS hashed_password VARCHAR not null default '123145123441fasd1!!4F1%&^';