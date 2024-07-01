CREATE TABLE "users"(
    "user_id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
    "user_name" VARCHAR NOT NULL,
    "user_email" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL,
    CONSTRAINT "users_pkey" PRIMARY KEY("user_id")
);

CREATE UNIQUE INDEX "users_name_key" ON "users"("user_name");