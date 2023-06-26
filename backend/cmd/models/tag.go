package models

type BoardTag struct {
	ID        int    `json:"id"`
	BoardID   int    `json:"board_id"`
	Title     string `json:"title"`
	Color     string `json:"color"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BoardTagUpdate struct {
	BoardID int    `json:"board_id"`
	Title   string `json:"title"`
	Color   string `json:"color"`
}

type TaskTag struct {
	TaskID     int    `json:"task_id"`
	BoardTagID int    `json:"board_tag_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
