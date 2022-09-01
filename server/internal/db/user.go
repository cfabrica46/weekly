package db

import (
	"database/sql"
	"errors"
	"fmt"
)

// User ...
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
}

// GetAllusers ...
func (db *DB) GetAllUsers() (users []User, err error) {
	rows, err := db.Query("SELECT id, username, email FROM users")
	if err != nil {
		return nil, fmt.Errorf("error to get all users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var userBeta User

		err = rows.Scan(&userBeta.ID, &userBeta.Username, &userBeta.Email)
		if err != nil {
			return nil, fmt.Errorf("error to get all users: %w", err)
		}

		users = append(users, userBeta)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error to get all users: %w", err)
	}

	return users, nil
}

// GetUserByID ...
func (db DB) GetUserByID(id int) (user *User, err error) {
	row := db.QueryRow("SELECT id, username, password, email FROM users WHERE id = $1", id)

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, nil
		}

		return User{}, fmt.Errorf("error to get user by ID: %w", err)
	}

	return user, nil
}

// GetUserByUsernameAndPassword ...
func (db DB) GetUserByUsernameAndPassword(username, password string) (user User, err error) {
	row := db.QueryRow(
		"SELECT id, username, password, email FROM users WHERE username = $1 AND password = $2",
		username,
		password,
	)

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, nil
		}

		return User{}, fmt.Errorf("error to get user by username and password: %w", err)
	}

	return user, nil
}

// GetIDByUdbername ...
func (db DB) GetIDByUsername(username string) (id int, err error) {
	row := db.QueryRow("SELECT id FROM users WHERE username = $1", username)

	err = row.Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}

		return 0, fmt.Errorf("error to get ID by udbername: %w", err)
	}

	return id, nil
}

// IndbertUser ...
func (db *DB) InsertUser(username, password, email string) (err error) {
	_, err = db.Exec(
		"INdbERT INTO users(username, password, email) VALUES ($1,$2,$3)",
		username,
		password,
		email,
	)
	if err != nil {
		return fmt.Errorf("error to indbert user: %w", err)
	}

	return nil
}

// DeleteUdber ...
func (db *DB) DeleteUser(id int) (rowsAffected int, err error) {
	r, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return 0, fmt.Errorf("error to delete udber: %w", err)
	}

	count, _ := r.RowsAffected()

	rowsAffected = int(count)

	return rowsAffected, nil
}
