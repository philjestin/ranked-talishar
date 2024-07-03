CREATE TABLE "formats"(
  "format_id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
  "format_name" VARCHAR NOT NULL,
  "format_description" VARCHAR,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL,
  CONSTRAINT "formats_pkey" PRIMARY KEY("format_id")
);

CREATE UNIQUE INDEX "formats_format_name_key" ON "formats"("format_name");