CREATE TABLE "board_column" (
    id SERIAL PRIMARY KEY,
    board_id INTEGER NOT NULL,
    name VARCHAR(30) NOT NULL,
    emoji VARCHAR(5) NOT NULL,
    position INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (board_id) REFERENCES board (id) ON DELETE CASCADE
);