CREATE TYPE room_type AS ENUM ('single', 'group');
CREATE TYPE message_type AS ENUM ('text', 'image', 'video', 'voice', 'file');
CREATE TYPE presence_status AS ENUM ('online', 'offline');

CREATE TABLE "rooms" (
  "id" UUID PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "type" room_type NOT NULL,
  "item_id" UUID,
  "project_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "messages" (
  "id" UUID PRIMARY KEY,
  "room_id" UUID NOT NULL REFERENCES "rooms"(id) ON DELETE CASCADE,
  "message" TEXT NOT NULL,
  "type" message_type NOT NULL DEFAULT 'text',
  "file" VARCHAR,
  "from" VARCHAR NOT NULL,
  "author_row_id" UUID NOT NULL,
  "read_at" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "room_members" (
  "id" UUID PRIMARY KEY,
  "room_id" UUID NOT NULL REFERENCES "rooms"(id) ON DELETE CASCADE,
  "row_id" UUID NOT NULL,
  "to_name" VARCHAR,
  "to_row_id" UUID,
  "last_read_at" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX ux_room_members_room_row ON "room_members"("room_id", "row_id");

CREATE TABLE "user_presence" (
  "row_id" UUID PRIMARY KEY,
  "status" presence_status NOT NULL DEFAULT 'online',
  "last_seen_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ix_user_presence_status ON "user_presence"("status");

CREATE INDEX ix_messages_room_created
  ON "messages"("room_id", "created_at" DESC);

CREATE INDEX ix_messages_room_unread
  ON "messages"("room_id")
  WHERE "read_at" IS NULL;

ALTER TABLE "messages"
ADD COLUMN parent_id UUID NULL REFERENCES "messages"(id) ON DELETE SET NULL;
