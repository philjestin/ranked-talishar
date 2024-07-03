CREATE TABLE "games"(
  "game_id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
  "game_name" VARCHAR NOT NULL,
  CONSTRAINT "games_pkey" PRIMARY KEY("game_id")
);

CREATE UNIQUE INDEX "games_name_key" ON "games"("game_name");