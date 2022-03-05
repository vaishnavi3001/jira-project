CREATE TABLE "users" (
	"user_id"	integer,
	"username"	text,
	"firstname"	text,
	"lastname"	text,
	"is_deleted" numeric NOT NULL DEFAULT 0,
	"email_id"	text,
	"created_at" datetime DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY("user_id" AUTOINCREMENT)
);

CREATE TABLE "projects" (
	"project_id"	integer,
	"project_name"	text NOT NULL,
	"is_deleted"	integer NOT NULL,
	"created_at"	datetime DEFAULT CURRENT_TIMESTAMP,
	"updated_at"	datetime,
	PRIMARY KEY("project_id" AUTOINCREMENT)
);

CREATE TABLE "issues" (
	"issue_id" integer,
	"status" integer DEFAULT 1,
	"type"	text,
	"title"	text,
	"created_by" integer, 
	"sprint_id"	text,
	"project_id" integer,
	"assignee_id" integer,
	"is_deleted" integer DEFAULT 0,
	"updated_at" datetime,
	"created_at" datetime DEFAULT CURRENT_TIMESTAMP,
	"description" text,
	PRIMARY KEY("issue_id" AUTOINCREMENT),
	FOREIGN KEY("sprint_id") REFERENCES "sprints"("sprint_id"),
	FOREIGN KEY("project_id") REFERENCES "projects"("project_id")
    FOREIGN KEY("created_by") REFERENCES "users"("user_id"),
    FOREIGN KEY("assignee_id") REFERENCES "users"("user_id")
)

CREATE TABLE "sprints" (
	"sprint_id"	integer,
	"sprint_name"	text,
	"project_id"	integer,
	"created_at"	datetime DEFAULT CURRENT_TIMESTAMP,
	"updated_at"	datetime,
	"start_date"	datetime,
	"end_date"	datetime,
	"status"	integer DEFAULT 1,
	PRIMARY KEY("sprint_id" AUTOINCREMENT),
	FOREIGN KEY("project_id") REFERENCES "projects"("project_id")
)

CREATE TABLE roles(
	role_id integer NOT NULL,
	role_name text NOT NULL,
	permission_id	integer,
	PRIMARY KEY(role_id),
	FOREIGN KEY(permission_id) REFERENCES permissions(permission_id)
)

CREATE TABLE user_roles(
    user_id integer,
    role_id integer, 
    project_id integer,
    membership integer,
    PRIMARY KEY (user_id, role_id, project_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (project_id) REFERENCES projects(project_id),
    FOREIGN KEY (role_id) REFERENCES roles(role_id)
)

CREATE TABLE "permissions" (
	"permission_id"	integer,
	"permission_name" text,
	PRIMARY KEY("permission_id")
)


