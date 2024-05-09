package users

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var dbString string = "root:strong_password@tcp(127.0.0.1:3306)/use_me_db"

type User struct {
	Id           int64  `json:"id"`
	Username     string `json:"username" gorm:"index;size:256"`
	PasswordHash string `binding:"required" gorm:"size:256"`
	Email        string `json:"email" gorm:"size:256"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func (u User) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a password to a hash and returns if it is valid or not.
func (u User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

func (u *User) SetPassword(password string) error {
	hash, err := u.hashPassword(password)

	if err == nil {
		u.PasswordHash = hash
	}

	return err
}

func (u *User) Save() error {
	db, err := sql.Open("mysql", dbString)
	if err == nil {
		defer db.Close()
		if insert, err := db.ExecContext(
			context.Background(), fmt.Sprintf(`
			INSERT INTO users(username, email, password_hash, created_at, updated_at)
			VALUE('%s', '%s', '%s', NOW(), NOW())`, u.Username, u.Email, u.PasswordHash)); err == nil {
			if id, err := insert.LastInsertId(); err == nil {
				if results, err := db.Query(fmt.Sprintf(`
					SELECT id, username, email, created_at, updated_at 
					FROM users 
					WHERE id=%v AND deleted_at IS NULL
					LIMIT 1`, id)); err == nil {
					results.Next()
					return results.Scan(&u.Id, &u.Username, &u.Email, &u.CreatedAt, &u.UpdatedAt) // scan error or nil
				}
				return err // db query error
			}
			return err // lastinsertId error
		}
		return err //insert error
	}
	return err //sql error
}

func FindByUsername(username string) User {
	fmt.Println("findByUsername : " + username)
	var u User
	db, err := sql.Open("mysql", dbString)
	if err == nil {
		defer db.Close()
		if results, err := db.Query(fmt.Sprintf(`
			SELECT id, username, email, password_hash, created_at, updated_at
			FROM users
			WHERE username = '%s' AND deleted_at IS NULL
			LIMIT 1
		`, username)); err == nil {
			results.Next()
			results.Scan(&u.Id, &u.Username, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt)
		}
	}
	return u
}
