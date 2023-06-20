CREATE TABLE check_test (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL
) 

CREATE TYPE "UserRole" AS ENUM ('ADMIN', 'USER');

CREATE TABLE user_table (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    inactive_status BOOLEAN NOT NULL DEFAULT FALSE,
    role "UserRole" NOT NULL DEFAULT 'USER'
);

CREATE TABLE "user_settings" (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    primary_board_id INTEGER,
    app_name TEXT NOT NULL,
    app_emoji TEXT NOT NULL,
    FOREIGN KEY (primary_board_id) REFERENCES board (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user_table (id) ON DELETE CASCADE
);

CREATE TABLE "board" (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    emoji TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user_table (id) ON DELETE CASCADE
);