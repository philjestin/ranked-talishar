CREATE TABLE "matches"(
  "match_id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
  "game_id" UUID REFERENCES games(game_id),
  "match_name" VARCHAR DEFAULT NULL,
  "player1_id" UUID REFERENCES users(user_id),
  "player2_id" UUID REFERENCES users(user_id),
  "winner_id" UUID REFERENCES users(user_id) DEFAULT NULL,
  "loser_id" UUID REFERENCES users(user_id) DEFAULT NULL,
  "player1_decklist" VARCHAR DEFAULT NULL,
  "player2_decklist" VARCHAR DEFAULT NULL,
  "player1_hero" UUID REFERENCES heroes(hero_id) DEFAULT NULL,
  "player2_hero" UUID REFERENCES heroes(hero_id) DEFAULT NULL,
  "match_date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "format_id" UUID REFERENCES formats(format_id) DEFAULT NULL,
  CONSTRAINT "matches_pkey" PRIMARY KEY("match_id")
);

CREATE UNIQUE INDEX "matches_match_id" on "matches"("match_id");