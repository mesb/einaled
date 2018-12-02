package user

import (
	"time"
)

// Name type for expressing the username of a user object
type Name string

// Service expresses the functions of the user package
//  the main crud procedures
// 1. add
// 2. getAll
// 3. getOne
// 4. delete
type Service interface {
	GetAll() ([]User, error)
	Lookup(Name) (*User, error)
	Put(User) error
}

// SecurityTest handles a secret question and answer for a user
// for password recovery or relevant security tests
type SecurityTest interface {
	question() string
	answer() string
}

// User captures the properties of a user on this platform
type User struct {
	Uname         Name
	Email         string
	password      string
	tests         []SecurityTest
	active        bool
	notes         string
	DateJoined    time.Time
	DatesModified []time.Time
}
