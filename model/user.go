package model

import "time"

type User struct {

	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at" pg:"default:now()"`
	UpdatedAt   time.Time `json:"updated_at"`
}
