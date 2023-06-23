package services

import (
	"time"

	"github.com/Dmkk01/kanban-go-svelte/cmd/db"
	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
)

func getSubTasksForTask(task *models.Task) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	taskID := task.ID

	rows, err := db.Query("SELECT id, task_id, title, completed, created_at, updated_at FROM subtasks WHERE task_id = $1", taskID)
	if err != nil {
		return err
	}

	task.SubTasks = []models.SubTask{}

	for rows.Next() {
		var subTask models.SubTask
		err := rows.Scan(&subTask.ID, &subTask.TaskID, &subTask.Title, &subTask.Completed, &subTask.CreatedAt, &subTask.UpdatedAt)
		if err != nil {
			return err
		}
		task.SubTasks = append(task.SubTasks, subTask)
	}
	defer db.Close()
	return nil
}

func getLinksForTask(task *models.Task) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	taskID := task.ID

	rows, err := db.Query("SELECT id, task_id, title, emoji, url, created_at, updated_at FROM links WHERE task_id = $1", taskID)
	if err != nil {
		return err
	}

	task.Links = []models.LinkTask{}

	for rows.Next() {
		var link models.LinkTask
		err := rows.Scan(&link.ID, &link.TaskID, &link.Title, &link.Emoji, &link.URL, &link.CreatedAt, &link.UpdatedAt)
		if err != nil {
			return err
		}
		task.Links = append(task.Links, link)
	}
	defer db.Close()
	return nil
}

func GetTaskByID(taskID int) (models.Task, error) {
	db, err := db.Connect()
	if err != nil {
		return models.Task{}, err
	}
	defer db.Close()

	var task models.Task
	err = db.QueryRow("SELECT id, board_id, column_id, title, description, position, time_needed, due_date, created_at, updated_at FROM task WHERE id = $1", taskID).Scan(&task.ID, &task.BoardID, &task.ColumnID, &task.Title, &task.Description, &task.Position, &task.TimeNeeded, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return models.Task{}, err
	}

	err = getSubTasksForTask(&task)
	if err != nil {
		return models.Task{}, err
	}

	err = getLinksForTask(&task)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func GetTaskByColumnID(columnID int) ([]models.Task, error) {
	db, err := db.Connect()
	if err != nil {
		return []models.Task{}, err
	}
	defer db.Close()

	tasks := []models.Task{}
	rows, err := db.Query("SELECT id, board_id, column_id, title, description, position, time_needed, due_date, created_at, updated_at FROM task WHERE column_id = $1 ORDER BY position ASC", columnID)
	if err != nil {
		return []models.Task{}, err
	}

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.BoardID, &task.ColumnID, &task.Title, &task.Description, &task.Position, &task.TimeNeeded, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return []models.Task{}, err
		}
		err = getSubTasksForTask(&task)
		if err != nil {
			return []models.Task{}, err
		}

		err = getLinksForTask(&task)
		if err != nil {
			return []models.Task{}, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func GetTaskByBoardID(boardID int) ([]models.Task, error) {
	db, err := db.Connect()
	if err != nil {
		return []models.Task{}, err
	}
	defer db.Close()

	tasks := []models.Task{}
	rows, err := db.Query("SELECT id, board_id, column_id, title, description, position, time_needed, due_date, created_at, updated_at FROM task WHERE board_id = $1", boardID)
	if err != nil {
		return []models.Task{}, err
	}

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.BoardID, &task.ColumnID, &task.Title, &task.Description, &task.Position, &task.TimeNeeded, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return []models.Task{}, err
		}
		err = getSubTasksForTask(&task)
		if err != nil {
			return []models.Task{}, err
		}

		err = getLinksForTask(&task)
		if err != nil {
			return []models.Task{}, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func GetAllTasks(userID int) ([]models.Task, error) {
	db, err := db.Connect()
	if err != nil {
		return []models.Task{}, err
	}
	defer db.Close()

	tasks := []models.Task{}
	rows, err := db.Query("SELECT task.id, task.board_id, task.column_id, task.title, task.description, task.position, task.time_needed, task.due_date, task.created_at, task.updated_at FROM task INNER JOIN board ON task.board_id = board.id WHERE board.user_id = $1", userID)
	if err != nil {
		return []models.Task{}, err
	}

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.BoardID, &task.ColumnID, &task.Title, &task.Description, &task.Position, &task.TimeNeeded, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return []models.Task{}, err
		}
		err = getSubTasksForTask(&task)
		if err != nil {
			return []models.Task{}, err
		}

		err = getLinksForTask(&task)
		if err != nil {
			return []models.Task{}, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func CreateTaskByColumnID(task models.TaskCreate) (int, error) {
	db, err := db.Connect()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	createdAt := time.Now()
	updatedAt := time.Now()

	taskID := 0
	err = db.QueryRow("INSERT INTO task (board_id, column_id, title, description, position, time_needed, due_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", task.BoardID, task.ColumnID, task.Title, task.Description, task.Position, task.TimeNeeded, task.DueDate, createdAt, updatedAt).Scan(&taskID)
	if err != nil {
		return 0, err
	}

	for _, subTask := range task.SubTasks {
		err := CreateSubTask(taskID, subTask)
		if err != nil {
			return 0, err
		}
	}

	for _, link := range task.Links {
		_, err = db.Exec("INSERT INTO links (task_id, title, emoji, url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", taskID, link.Title, link.Emoji, link.URL, createdAt, updatedAt)
		if err != nil {
			return 0, err
		}
	}

	return int(taskID), nil
}

func DeleteTaskByID(taskID int) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM task WHERE id = $1", taskID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTaskByID(taskID int, task models.TaskUpdate) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	updatedAt := time.Now()

	_, err = db.Exec("UPDATE task SET title = $1, description = $2, position = $3, time_needed = $4, due_date = $5, updated_at = $6 WHERE id = $7", task.Title, task.Description, task.Position, task.TimeNeeded, task.DueDate, updatedAt, taskID)
	if err != nil {
		return err
	}

	for _, subTask := range task.SubTasks {
		if subTask.ID == 0 {
			err := CreateSubTask(taskID, models.SubTaskCreate{
				Title:     subTask.Title,
				Completed: subTask.Completed,
			})
			if err != nil {
				return err
			}
		} else {
			err := UpdateSubTaskByID(models.SubTaskUpdate{
				ID:        subTask.ID,
				Title:     subTask.Title,
				Completed: subTask.Completed,
			})
			if err != nil {
				return err
			}
		}
	}

	for _, link := range task.Links {
		if link.ID == 0 {
			_, err = db.Exec("INSERT INTO links (task_id, title, emoji, url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", taskID, link.Title, link.Emoji, link.URL, updatedAt, updatedAt)
			if err != nil {
				return err
			}
		} else {
			_, err = db.Exec("UPDATE links SET title = $1, emoji = $2, url = $3, updated_at = $4 WHERE id = $5", link.Title, link.Emoji, link.URL, updatedAt, link.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func UpdateTaskPosition(pos models.TaskPositionUpdate) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	updatedAt := time.Now()
	_, err = db.Exec("UPDATE task SET position = $1, updated_at = $2, column_id = $3 WHERE id = $4", pos.Position, updatedAt, pos.ColumnID, pos.TaskID)

	return err
}

func GetSubTasks(taskID int) ([]models.SubTask, error) {
	db, err := db.Connect()
	if err != nil {
		return []models.SubTask{}, err
	}
	defer db.Close()

	subTasks := []models.SubTask{}
	rows, err := db.Query("SELECT id, task_id, title, completed, created_at, updated_at FROM subtasks WHERE task_id = $1", taskID)
	if err != nil {
		return []models.SubTask{}, err
	}

	for rows.Next() {
		var subTask models.SubTask
		err := rows.Scan(&subTask.ID, &subTask.TaskID, &subTask.Title, &subTask.Completed, &subTask.CreatedAt, &subTask.UpdatedAt)
		if err != nil {
			return []models.SubTask{}, err
		}

		subTasks = append(subTasks, subTask)
	}

	return subTasks, nil
}

func GetSubTaskByID(subtaskID int) (models.SubTask, error) {
	db, err := db.Connect()
	if err != nil {
		return models.SubTask{}, err
	}
	defer db.Close()

	var subTask models.SubTask
	err = db.QueryRow("SELECT id, task_id, title, completed, created_at, updated_at FROM subtasks WHERE id = $1", subtaskID).Scan(&subTask.ID, &subTask.TaskID, &subTask.Title, &subTask.Completed, &subTask.CreatedAt, &subTask.UpdatedAt)
	if err != nil {
		return models.SubTask{}, err
	}

	return subTask, nil
}

func CreateSubTask(taskID int, subtask models.SubTaskCreate) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	createdAt := time.Now()
	_, err = db.Exec("INSERT INTO subtasks (task_id, title, completed, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", taskID, subtask.Title, subtask.Completed, createdAt, createdAt)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSubTaskByID(subtaskID int) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM subtasks WHERE id = $1", subtaskID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateSubTaskByID(subtask models.SubTaskUpdate) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	updatedAt := time.Now()
	_, err = db.Exec("UPDATE subtasks SET title = $1, completed = $2, updated_at = $3 WHERE id = $4", subtask.Title, subtask.Completed, updatedAt, subtask.ID)

	return err
}
