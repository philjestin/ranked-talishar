CREATE TABLE rankings (
  "ranking_id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
  "user_id" UUID REFERENCES users(user_id),
  "game_id" UUID REFERENCES games(game_id),
  "format_id" UUID REFERENCES formats(format_id),
  "elo_score" INTEGER,
  "last_updated" TIMESTAMP DEFAULT NULL,
  CONSTRAINT "ranking_pkey" PRIMARY KEY("ranking_id")
);

CREATE UNIQUE INDEX "rankings_ranking_id_key" ON "rankings"("ranking_id");