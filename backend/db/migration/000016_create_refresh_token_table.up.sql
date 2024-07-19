CREATE TABLE "refresh_tokens" (
  "id" SERIAL PRIMARY key,
  "user_id" UUID NOT NULL,
  "refresh_token" VARCHAR NOT NULL,  
  "expiry" TIMESTAMP NOT NULL,
  FOREIGN KEY ("user_id") REFERENCES users(user_id)
);