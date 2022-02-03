CREATE TABLE IF NOT EXISTS "todos"
  (
     "id"           BIGSERIAL,
     "created_at"   TIMESTAMPTZ NOT NULL,
     "updated_at"   TIMESTAMPTZ NOT NULL,
     "deleted_at"   TIMESTAMPTZ,
     "content_code" TEXT NOT NULL UNIQUE,
     "content_name" TEXT,
     "description"  TEXT,
     "start_date"   TIMESTAMPTZ,
     "end_date"     TIMESTAMPTZ,
     "status"       TEXT,
     "created_by"   TEXT,
     "updated_by"   TEXT,

     PRIMARY KEY ("id")
  );

CREATE INDEX "idx_todo_deleted_at" ON "todos" ("deleted_at")
