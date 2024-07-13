ALTER TABLE users
ADD column if not exists password_changed_at TIMESTAMP;