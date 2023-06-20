CREATE TABLE "user_settings" (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    primary_board_id INTEGER,
    app_name TEXT NOT NULL,
    app_emoji TEXT NOT NULL,
    FOREIGN KEY (primary_board_id) REFERENCES board (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user_table (id) ON DELETE CASCADE
);