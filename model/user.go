package model

import "time"

type User struct {
	tableName struct{}  `pg:"auth.users"`               //bảng users có schema auth
	Id        int       `json:"id" pg:"type:serial,pk"` //trường id primary key, auto_increment.
	FirstName string    `json:"first_name" pg:"type:varchar(50),notnull"`             //trường first_name text
	LastName  string    `json:"last_name" pg:"type:varchar(50),notnull"`              //trường last_name  text
	Email     string    `json:"email" pg:"type:varchar(255),unique"`     //trường email không được trùng lặp
	Password  string    `json:"password"`               //trường password text
	CreatedAt time.Time `json:"created_at" pg:"type:timestamp without time zone,default:now()"`
	UpdatedAt time.Time `json:"updated_at" pg:"type:timestamp without time zone"`

	UserRoles []UserRole `json:"user_roles" pg:"rel:has-many"`
}
