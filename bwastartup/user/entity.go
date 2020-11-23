package user

import "time"

type User struct {
	ID               int
	Name             string
	Occupation       string
	Email            string
	Password_hash    string
	Avatar_file_name string
	Role             string
	Token            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
