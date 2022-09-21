package db

import (
	"fmt"
)

type DayOfWeek string

type DaysOfWeek struct {
	Monday    bool
	Tuesday   bool
	Wednesday bool
	Thursday  bool
	Friday    bool
	Saturday  bool
	Sunday    bool
}

// Task ...
type Task struct {
	Title       string
	Description string
	ID          int
	Days        DaysOfWeek
}

// Tasks ...
type Tasks []Task

const (
	Monday    DayOfWeek = "monday"
	Tuesday   DayOfWeek = "tuesday"
	Wednesday DayOfWeek = "wednesday"
	Thursday  DayOfWeek = "thursday"
	Friday    DayOfWeek = "friday"
	Saturday  DayOfWeek = "saturday"
	Sunday    DayOfWeek = "sunday"
)

const (
	queryGetAllTasks = "SELECT id, title, description, monday, tuesday, wednesday, thursday, friday, saturday, sunday FROM tasks"

	queryGetOneDayTasks = "SELECT id, title, description, monday, tuesday, wednesday, thursday, friday, saturday, sunday FROM tasks WHERE %s = true"

	queryInsertTask = "INSERT INTO tasks(title, description, monday, tuesday, wednesday, thursday, friday, saturday, sunday) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)"

	queryToggleDayOfWeek = "UPDATE tasks SET %s = NOT %s WHERE id = $1"
)

func (d DaysOfWeek) String() string {
	var monday, tuesday, wednesday, thursday, friday, saturday, sunday string

	if d.Monday {
		monday = "monday "
	}

	if d.Tuesday {
		tuesday = "tuesday "
	}

	if d.Wednesday {
		wednesday = "wednesday "
	}

	if d.Thursday {
		thursday = "thursday "
	}

	if d.Friday {
		friday = "friday "
	}

	if d.Saturday {
		saturday = "saturday "
	}

	if d.Sunday {
		sunday = "sunday "
	}

	return fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		monday,
		tuesday,
		wednesday,
		thursday,
		friday,
		saturday,
		sunday,
	)
}

func (t Task) String() string {
	return fmt.Sprintf(
		"ID: %d | Title: %s | Description: %s | Days: %s",
		t.ID,
		t.Title,
		t.Description,
		fmt.Sprint(t.Days),
	)
}

func (t Tasks) String() string {
	var str string

	if t == nil {
		return "[]"
	}

	str = "ID\tTITLE\tDESCRIPTION\tDAYS\n--\t-----\t-----------\t----\n"

	for i := range t {
		str += fmt.Sprintf("%d\t%s\t%s\t%s\n", t[i].ID, t[i].Title, t[i].Description, fmt.Sprint(t[i].Days))
	}

	return str
}

// GetAllTasks ...
func (db *DB) GetAllTasks() (tasks Tasks, err error) {
	return GetTasks(db, queryGetAllTasks)
}

// GetOneDayTask ...
func (db *DB) GetOneDayTask(day DayOfWeek) (tasks Tasks, err error) {
	return GetTasks(db, fmt.Sprintf(queryGetOneDayTasks, day))
}

func GetTasks(db *DB, query string, args ...any) (tasks Tasks, err error) {
	rows, err := db.Query(query, args...)
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
			&taskBeta.Days.Monday,
			&taskBeta.Days.Tuesday,
			&taskBeta.Days.Wednesday,
			&taskBeta.Days.Thursday,
			&taskBeta.Days.Friday,
			&taskBeta.Days.Saturday,
			&taskBeta.Days.Sunday,
		)
		if err != nil {
			return nil, fmt.Errorf("error to get tasks: %w", err)
		}

		tasks = append(tasks, taskBeta)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error to get users: %w", err)
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
		task.Days.Monday,
		task.Days.Tuesday,
		task.Days.Wednesday,
		task.Days.Thursday,
		task.Days.Friday,
		task.Days.Saturday,
		task.Days.Sunday,
	)
	if err != nil {
		return fmt.Errorf("error to insert task: %w", err)
	}

	return nil
}

// ToggleDayOfTask ...
func (db *DB) ToggleDayOfTask(id int, day DayOfWeek) (err error) {
	query := fmt.Sprintf(queryToggleDayOfWeek, day, day)

	_, err = db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error to Toggle Day of Week task: %w", err)
	}

	return nil
}
