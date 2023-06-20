package models

type BoardColumn struct {
	ID        int    `json:"id"`
	BoardID   int    `json:"board_id"`
	Name      string `json:"name"`
	Emoji     string `json:"emoji"`
	Position  int    `json:"position"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ColumnCreate struct {
	Name     string `json:"name"`
	Emoji    string `json:"emoji"`
	Position int    `json:"position"`
}
