package db

import "fmt"

type DayOfWeek string

// Task ...
type Task struct {
	Title       string
	Description string
	ID          int
	Monday      bool
	Tuesday     bool
	Wednesday   bool
	Thursday    bool
	Friday      bool
	Saturday    bool
	Sunday      bool
}

const (
	queryGetAllTasks = "SELECT id, title, description, monday, tuesday, wednesday, thursday, friday, saturday, sunday FROM tasks"

	queryGetOneDayTasks = "SELECT id, title, description, monday, tuesday, wednesday, thursday, friday, saturday, sunday FROM tasks WHERE $1 = true"

	queryInsertTask = "INSERT INTO tasks(title, description, monday, tuesday, wednesday, thursday, friday, saturday, sunday) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)"

	queryToggleDayOfWeek = "UPDATE tasks SET $1 = NOT $1 WHERE id = $2"
)

// GetAllTasks ...
func (db *DB) GetAllTasks() (tasks []Task, err error) {
	rows, err := db.Query(queryGetAllTasks)
	if err != nil {
		return nil, fmt.Errorf("error when creating the query : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var taskBeta Task

		err = rows.Scan(
			&taskBeta.ID,
			&taskBeta.Title,
			&taskBeta.Description,
			&taskBeta.Monday,
			&taskBeta.Tuesday,
			&taskBeta.Wednesday,
			&taskBeta.Thursday,
			&taskBeta.Friday,
			&taskBeta.Saturday,
			&taskBeta.Sunday,
		)
		if err != nil {
			return nil, fmt.Errorf("error to get all tasks: %w", err)
		}

		tasks = append(tasks, taskBeta)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error to get all users: %w", err)
	}

	return tasks, nil
}

// GetOneDayTask ...
func (db *DB) GetOneDayTask() (tasks []Task, err error) {
	rows, err := db.Query(queryGetOneDayTasks)
	if err != nil {
		return nil, fmt.Errorf("error when creating the query : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var taskBeta Task

		err = rows.Scan(
			&taskBeta.ID,
			&taskBeta.Title,
			&taskBeta.Description,
			&taskBeta.Monday,
			&taskBeta.Tuesday,
			&taskBeta.Wednesday,
			&taskBeta.Thursday,
			&taskBeta.Friday,
			&taskBeta.Saturday,
			&taskBeta.Sunday,
		)
		if err != nil {
			return nil, fmt.Errorf("error to get all tasks: %w", err)
		}

		tasks = append(tasks, taskBeta)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error to get all users: %w", err)
	}

	return tasks, nil
}

// InsertTask ...
func (db *DB) InsertTask(task Task) (err error) {
	_, err = db.Exec(
		queryInsertTask,
		task.ID,
		task.Title,
		task.Description,
		task.Monday,
		task.Tuesday,
		task.Wednesday,
		task.Thursday,
		task.Friday,
		task.Saturday,
		task.Sunday,
	)
	if err != nil {
		return fmt.Errorf("error to insert task: %w", err)
	}

	return nil
}

// ToggleDayOfTask ...
func (db *DB) ToggleDayOfTask(id int, day DayOfWeek) (err error) {
	_, err = db.Exec(
		queryToggleDayOfWeek,
		day,
		id,
	)
	if err != nil {
		return fmt.Errorf("error to Toggle Day of Week task: %w", err)
	}

	return nil
}
