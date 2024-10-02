package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type UserModel struct {
	DB *sql.DB
}

func (db *UserModel) Insert(user *User) error {
	stmt := `
    INSERT INTO Users (username, email)
    VALUES ($1, $2) 
    RETURNING id;`
	args := []interface{}{user.UserName, user.Email}

	err := db.Find(user)
	if err == nil {
		return errors.New("the given element is already taken")
	}

	if err != sql.ErrNoRows {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = db.DB.QueryRowContext(ctx, stmt, args...).Scan(&user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (db *UserModel) Find(user *User) error {
	stmt := `
    SELECT id, username, email
    FROM Users
    WHERE
      id = $1
  `
	args := []interface{}{user.Id}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := db.DB.QueryRowContext(ctx, stmt, args...).Scan(&user.Id, &user.UserName, &user.Email)
	if err != nil {
		return err
	}

	return nil
}
