package model

import "time"

type UserRole struct {
	tableName struct{}  `pg:"auth.user_role"`
	UserId    int       `json:"user_id" pg:"type:integer"`
	RoleId    int       `json:"role_id" pg:"type:integer"`
	User      User      `pg:"rel:has-one"`
	Role      Role      `pg:"rel:one-one"`
	CreatedAt time.Time `json:"created_at" pg:"type:timestamp without time zone,default:now()"`
	UpdatedAt time.Time `json:"updated_at" pg:"type:timestamp without time zone"`
}
