-- Adminer 4.8.1 PostgreSQL 16.6 (Debian 16.6-1.pgdg120+1) dump

DROP TABLE IF EXISTS "cronjobs";
DROP SEQUENCE IF EXISTS cronjob_id_seq;
CREATE SEQUENCE cronjob_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."cronjobs" (
    "id" bigint DEFAULT nextval('cronjob_id_seq') NOT NULL,
    "name" character varying(255) NOT NULL,
    "schedule" character varying NOT NULL,
    "task" character varying NOT NULL,
    "status" boolean NOT NULL,
    "user_id" bigint NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" timestamp NOT NULL,
    CONSTRAINT "cronjob_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "logs";
DROP SEQUENCE IF EXISTS logs_id_seq;
CREATE SEQUENCE logs_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."logs" (
    "id" bigint DEFAULT nextval('logs_id_seq') NOT NULL,
    "job_id" bigint NOT NULL,
    "execution_time" timestamp NOT NULL,
    "status" boolean NOT NULL,
    "duration" integer NOT NULL,
    "error_message" text NOT NULL,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "updated_at" timestamp NOT NULL,
    CONSTRAINT "logs_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "users";
DROP SEQUENCE IF EXISTS user_id_seq;
CREATE SEQUENCE user_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."users" (
    "id" bigint DEFAULT nextval('user_id_seq') NOT NULL,
    "name" character varying NOT NULL,
    "email" character varying NOT NULL,
    "password" character varying NOT NULL,
    "created_at" timestamp DEFAULT now(),
    "updated_at" timestamp NOT NULL,
    CONSTRAINT "user_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


-- 2025-01-28 13:47:53.064006+00
