CREATE TABLE "Users"(
    "id" UUID NOT NULL,
    "nickname" VARCHAR(50) NOT NULL DEFAULT 'guest{id}',
    "avatar" TEXT NOT NULL,
    "email" VARCHAR(100) NOT NULL,
    "password_hash" VARCHAR(255) NOT NULL,
    "oath_provider" VARCHAR(255) NULL,
    "id_oath" BIGINT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "Users" ADD PRIMARY KEY("id");
ALTER TABLE
    "Users" ADD CONSTRAINT "users_email_unique" UNIQUE("email");
CREATE TABLE "Sessions"(
    "id" UUID NOT NULL,
    "id_owner" UUID NULL,
    "title" VARCHAR(50) NULL,
    "language" VARCHAR(50) NOT NULL,
    "access_type" VARCHAR(255) CHECK("access_type" IN ('Public', 'Private')) NOT NULL DEFAULT 'Public',
        "expiration_time" TIME(0) WITHOUT TIME ZONE NULL,
        "max_users" BIGINT NOT NULL DEFAULT '20',
        "is_editable" BOOLEAN NOT NULL DEFAULT '1',
        "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
        "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
        "is_active" BOOLEAN NOT NULL DEFAULT '0'
);
ALTER TABLE
    "Sessions" ADD PRIMARY KEY("id");
CREATE TABLE "CodeSnippets"(
    "id" UUID NOT NULL,
    "id_session" UUID NOT NULL,
    "language" VARCHAR(50) NOT NULL,
    "code" TEXT NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "CodeSnippets" ADD PRIMARY KEY("id");
CREATE TABLE "Comments"(
    "id" UUID NOT NULL,
    "id_session" UUID NOT NULL,
    "id_author" UUID NULL,
    "content" TEXT NOT NULL,
    "line_start" BIGINT NOT NULL,
    "line_end" BIGINT NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "Comments" ADD PRIMARY KEY("id");
CREATE TABLE "SessionParticipants"(
    "id" UUID NOT NULL,
    "id_session" UUID NOT NULL,
    "id_user" UUID UNIQUE NULL,
    "nickname" VARCHAR(50) NOT NULL,
    "avatar" TEXT NOT NULL,
    "is_creator" BOOLEAN NOT NULL,
    "can_edit" BOOLEAN NOT NULL,
    "joined_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "SessionParticipants" ADD PRIMARY KEY("id");
CREATE TABLE "CodeVersions"(
    "id" UUID NOT NULL,
    "id_snippet" UUID NOT NULL,
    "code" TEXT NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "id_author" UUID NULL
);
ALTER TABLE
    "CodeVersions" ADD PRIMARY KEY("id");
CREATE TABLE "Templates"(
    "id" BIGINT NOT NULL,
    "name" VARCHAR(50) NOT NULL,
    "language" VARCHAR(50) NOT NULL,
    "template_code" TEXT NOT NULL,
    "created_by" UUID NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "Templates" ADD PRIMARY KEY("id");
CREATE TABLE "Messages"(
    "id" BIGINT NOT NULL,
    "id_session" UUID UNIQUE NOT NULL,
    "id_user" UUID NOT NULL,
    "message" TEXT NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "Messages" ADD PRIMARY KEY("id");
ALTER TABLE
    "SessionParticipants" ADD CONSTRAINT "sessionparticipants_id_user_foreign" FOREIGN KEY("id_user") REFERENCES "Users"("id");
ALTER TABLE
    "Messages" ADD CONSTRAINT "messages_id_user_foreign" FOREIGN KEY("id_user") REFERENCES "SessionParticipants"("id_user");
ALTER TABLE
    "SessionParticipants" ADD CONSTRAINT "sessionparticipants_id_session_foreign" FOREIGN KEY("id_session") REFERENCES "Sessions"("id");
ALTER TABLE
    "CodeVersions" ADD CONSTRAINT "codeversions_id_snippet_foreign" FOREIGN KEY("id_snippet") REFERENCES "CodeSnippets"("id");
ALTER TABLE
    "Comments" ADD CONSTRAINT "comments_id_session_foreign" FOREIGN KEY("id_session") REFERENCES "Sessions"("id");
ALTER TABLE
    "CodeSnippets" ADD CONSTRAINT "codesnippets_id_session_foreign" FOREIGN KEY("id_session") REFERENCES "Sessions"("id");
ALTER TABLE
    "Templates" ADD CONSTRAINT "templates_created_by_foreign" FOREIGN KEY("created_by") REFERENCES "Users"("id");
