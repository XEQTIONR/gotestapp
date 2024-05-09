package user

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	passwordHash string `binding:"required"`
	Email        string `json:"email"`
}

func (u User) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a password to a hash and returns if it is valid or not.
func (u User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.passwordHash), []byte(password))
	return err == nil
}

func (u *User) SetPassword(password string) error {
	hash, err := u.hashPassword(password)

	if err == nil {
		u.passwordHash = hash
	}

	return err
}
