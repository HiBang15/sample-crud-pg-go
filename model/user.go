package model

import "time"

type User struct {
	tableName struct{} `pg:"auth.users"` //bảng users có schema auth
	Id int `pg:"type:serial,pk"` //trường id primary key, auto_increment.
	FirstName string //trường first_name text
	LastName string //trường last_name  text
	Email string `pg:",unique"` //trường email không được trùng lặp
	Password  string    //trường password text
	CreatedAt time.Time `json:"created_at" pg:"default:now()"`
	UpdatedAt time.Time `json:"updated_at"`
}
