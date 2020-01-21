/**
 * @author Mory Keita on 1/20/20
 */
package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)
type UserID string

type User struct {
	ID           UserID       `json:"id,omitempty" db:"user_id"`
	Email        *string      `json:"email" db:"emil"`
	PasswordHash *[]byte      `json:"_" db:"password_hash"`
	CreateAt     *[]time.Time `json:"_" db:"created_at"`
	DeletedAt    *[]time.Time `json:"_" db:"deleted_at"`
}
// update user's password
func (u *User) SetPassword (password string) error{
	hash , err := HashPassword(password)
	if err != nil{
		return  err
	}
	u.PasswordHash = &hash
	return nil
}

// Verify all fields before create or update

func (u *User) Verify() error {
	if u.Email == nil || len(*u.Email) == 0{
		return errors.New("Email is required.")
	}
	return nil
}

func (u *User) checkPassword (password string) error  {
	if u.PasswordHash != nil && len(*u.PasswordHash) == 0 {
		return errors.New("Password not set.")
	}
	return bcrypt.CompareHashAndPassword(*u.PasswordHash,[]byte(password))
}
// Hash user's raw password using crypto
func HashPassword(password string) ([] byte,error) {
	return bcrypt.GenerateFromPassword([]byte (password),bcrypt.DefaultCost)
}
