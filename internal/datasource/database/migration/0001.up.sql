CREATE TABLE IF NOT EXISTS account (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "name" VARCHAR(255) NOT NULL,
    "initial_value" REAL NOT NULL,
    "current_value" REAL NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP
);
