package model

import "time"

type Role struct {
	tableName struct{}  `pg:"auth.roles"`
	Id        int       `json:"id" pg:"type:serial,pk"`
	Name      string    `json:"name" pg:"type:varchar(50),notnull"`
	CreatedAt time.Time `json:"created_at" pg:"type:timestamp without time zone,default:now()"`
	UpdatedAt time.Time `json:"updated_at" pg:"type:timestamp without time zone"`

	Users []User `json:"users" pg:"rel:has-many"`
	Permissions []Permission `json:"permissions" pg:"has-many"`
}
