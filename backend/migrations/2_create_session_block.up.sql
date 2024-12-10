CREATE TABLE "SessionBlock"(
    "id" UUID NOT NULL,
    "ip_address" INET NOT NULL,
    "session_id" UUID NOT NULL
);

ALTER TABLE
    "SessionBlock" ADD PRIMARY KEY("id");

ALTER TABLE
    "SessionBlock" ADD CONSTRAINT "sessionblock_session_id_foreign" FOREIGN KEY("session_id") REFERENCES "Sessions"("id");

CREATE INDEX IF NOT EXISTS "idx_ip" ON "SessionBlock"("ip_address");
