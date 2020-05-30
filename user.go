// Package terrorblade ...
package terrorblade

import "github.com/jinzhu/gorm"

// User represent User of system.
type User struct {
	ID       int64      `db:"id"`
	FullName string     `db:"full_name"`
	Username string     `db:"username"`
	Password string     `db:"password"`
	Status   UserStatus `db:"status"`
}

// TableName returns User's table name.
func (User) TableName() string {
	return "user"
}

// UserStatus represents status of a User.
type UserStatus string

// UserStatus definition.
const (
	UserStatusActive   = UserStatus("ACTIVE")
	UserStatusInActive = UserStatus("INACTIVE")
)

// UserRepository provides a access to User store.
type UserRepository interface {
	CreateUser(db *gorm.DB, user *User) error
	GetUser(db *gorm.DB, id int64) (*User, error)
}
