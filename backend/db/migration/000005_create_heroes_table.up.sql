CREATE TABLE "heroes"(
  "hero_id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
  "hero_name" VARCHAR NOT NULL,
  "format_id" UUID REFERENCES formats(format_id),
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL,
  CONSTRAINT "heroes_pkey" PRIMARY KEY("hero_id")
);

CREATE UNIQUE INDEX "heroes_hero_name_key" ON "heroes"("hero_name");