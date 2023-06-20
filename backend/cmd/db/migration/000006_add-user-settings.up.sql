DROP TABLE IF EXISTS "user_settings";
CREATE TABLE "user_settings" (
    user_id INTEGER NOT NULL,
    primary_board_id INTEGER,
    app_name VARCHAR(30) NOT NULL,
    app_emoji VARCHAR(5) NOT NULL,
    PRIMARY KEY (user_id),
    CONSTRAINT fk_user_settings_user_id FOREIGN KEY (user_id) REFERENCES user_table (id) ON DELETE CASCADE,
    FOREIGN KEY (primary_board_id) REFERENCES board (id) ON DELETE CASCADE
);