CREATE TABLE IF NOT EXISTS account (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "name" VARCHAR(255) NOT NULL,
    "initial_value" REAL NOT NULL,
    "current_value" REAL NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS account_transaction (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "account_id" INTEGER NOT NULL REFERENCES account (id) ON DELETE CASCADE,
    "name" VARCHAR(255) NOT NULL,
    "value" REAL NOT NULL,
    "created_at" TIMESTAMP NOT NULL
);
