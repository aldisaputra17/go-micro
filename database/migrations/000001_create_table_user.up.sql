CREATE TABLE IF NOT EXISTS "users" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL DEFAULT 'system',
  "updated_at" timestamp,
  "updated_by" varchar,
  "deleted_at" timestamp,
  "deleted_by" varchar
);

