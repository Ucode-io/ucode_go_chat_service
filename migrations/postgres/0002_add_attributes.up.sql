ALTER TABLE "rooms" ADD COLUMN "attributes" JSONB DEFAULT '{}'::jsonb;
ALTER TABLE "room_members" ADD COLUMN "attributes" JSONB DEFAULT '{}'::jsonb;

