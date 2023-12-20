CREATE TABLE IF NOT EXISTS "users" (
  "id" SERIAL PRIMARY KEY,
  "firstname" VARCHAR (64) NOT NULL,
  "lastname" VARCHAR (64) NOT NULL,
  "email" VARCHAR (128) UNIQUE NOT NULL,
  "status" BOOLEAN NOT NULL DEFAULT FALSE,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "posts" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INT NOT NULL,
  "title" VARCHAR (128) NOT NULL,
  "description" TEXT NOT NULL,
  "status" BOOLEAN NOT NULL DEFAULT FALSE,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "comments" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INT NOT NULL,
  "post_id" INT NOT NULL,
  "title" VARCHAR (128) NOT NULL,
  "description" TEXT NOT NULL,
  "status" BOOLEAN NOT NULL DEFAULT FALSE,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");
ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX ON "posts" ("user_id");
CREATE INDEX ON "comments" ("post_id");
CREATE INDEX ON "comments" ("user_id");