-- Modify "projects" table
ALTER TABLE "projects" ADD COLUMN "slot_uri" jsonb NOT NULL DEFAULT '{}'::jsonb;
