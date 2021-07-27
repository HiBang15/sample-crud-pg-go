package model

import "time"

type Permission struct {
	tableName struct{}  `pg:"auth.permission"`
	Id        int       `json:"id" pg:"type:serial,pk"`
	Name      string    `json:"name" pg:"type:varchar(50),notnull"`
	CreatedAt time.Time `json:"created_at" pg:"type:timestamp without time zone,default:now()"`
	UpdatedAt time.Time `json:"updated_at" pg:"type:timestamp without time zone"`

	Activities []Activity `json:"activities" pg:"rel:has-many"`
	Permissions []Permission `json:"permissions" pg:"rel:has-many"`
}
