package models

type Board struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Name      string `json:"name"`
	Emoji     string `json:"emoji"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
