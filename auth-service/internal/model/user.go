/**
 * @author Mory Keita on 1/20/20
 */
package model

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)
type UserID string

type User struct {
	ID UserID    'json:"id,omitempty" db:"user_id"'
	Eamil  *string    'json:"email" db:"emil"'
	PasswordHash *[]byte 'json:"_" db:"password_hash"'
	CreateAt *[]time.Time 'json:"_" db:"deleted_at"'
}



func (u *User) SetPassword (password string) error{
	return  nil
}
// Hash user's raw password using crypto
func HashPassword(password string) ([] byte,error) {
	hash, err := bcrypt.GenerateFromPassword([]byte (password),bcrypt.DefaultCost)
	if err != nil{

	}
	return hash,nil
}
