CREATE TABLE "task" (
  id VARCHAR PRIMARY KEY,
  title VARCHAR NOT NULL,
  description TEXT,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  limited_at TIMESTAMPTZ,
  is_end BOOLEAN DEFAULT FALSE
);
CREATE OR REPLACE FUNCTION update_updated_at_column() RETURNS TRIGGER AS $$ BEGIN NEW.updated_at = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER set_updated_at BEFORE
UPDATE ON "task" FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TABLE "tag" (
  id VARCHAR PRIMARY KEY,
  name VARCHAR NOT NULL
);
CREATE TABLE "task_tag" (
  task_id VARCHAR REFERENCES task(id) ON DELETE CASCADE,
  tag_id VARCHAR REFERENCES tag(id) ON DELETE CASCADE
);
ALTER TABLE "task_tag"
ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id");
ALTER TABLE "task_tag"
ADD FOREIGN KEY ("tag_id") REFERENCES "tag" ("id");