package db

import "fmt"

/* id INT GENERATED ALWAYS AS IDENTITY,
	user_id INT REFERENCES users,
    title VARCHAR(64) NOT NULL UNIQUE,
    description TEXT NOT NULL UNIQUE,
	monday BOOLEAN,
	tuesday BOOLEAN,
	wednesday BOOLEAN,
	thursday BOOLEAN,
	friday BOOLEAN,
	saturday BOOLEAN,
	sunday BOOLEAN, */

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
