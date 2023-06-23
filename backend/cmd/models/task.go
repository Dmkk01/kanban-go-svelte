package models

type Task struct {
	ID          int        `json:"id"`
	BoardID     int        `json:"board_id"`
	ColumnID    int        `json:"column_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	TimeNeeded  string     `json:"time_needed"`
	DueDate     string     `json:"due_date"`
	Position    int        `json:"position"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
	SubTasks    []SubTask  `json:"sub_tasks"`
	Links       []LinkTask `json:"links"`
}

type SubTask struct {
	ID        int    `json:"id"`
	TaskID    int    `json:"task_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type LinkTask struct {
	ID        int    `json:"id"`
	TaskID    int    `json:"task_id"`
	Title     string `json:"title"`
	Emoji     string `json:"emoji"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TaskCreate struct {
	BoardID     int              `json:"board_id"`
	ColumnID    int              `json:"column_id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	TimeNeeded  int              `json:"time_needed"`
	DueDate     string           `json:"due_date"`
	Position    int              `json:"position"`
	SubTasks    []SubTaskCreate  `json:"sub_tasks"`
	Links       []LinkTaskCreate `json:"links"`
}

type SubTaskCreate struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type LinkTaskCreate struct {
	Title string `json:"title"`
	Emoji string `json:"emoji"`
	URL   string `json:"url"`
}

type TaskPositionUpdate struct {
	TaskID   int `json:"task_id"`
	ColumnID int `json:"column_id"`
	Position int `json:"position"`
}

type TaskUpdate struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	TimeNeeded  int              `json:"time_needed"`
	DueDate     string           `json:"due_date"`
	Position    int              `json:"position"`
	SubTasks    []SubTaskUpdate  `json:"sub_tasks"`
	Links       []LinkTaskUpdate `json:"links"`
}

type SubTaskUpdate struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type LinkTaskUpdate struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Emoji string `json:"emoji"`
	URL   string `json:"url"`
}
