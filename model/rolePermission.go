package model

import "time"

type RolePermission struct {
	tableName    struct{}   `pg:"auth.role_permission"`
	RoleId       int        `json:"role_id" pg:"type:integer"`
	PermissionId int        `json:"permission_id" pg:"type:integer"`
	Role         Role       `pg:"rel:has-one"`
	Permission   Permission `pg:"rel:one-one"`
	CreatedAt    time.Time  `json:"created_at" pg:"type:timestamp without time zone,default:now()"`
	UpdatedAt    time.Time  `json:"updated_at" pg:"type:timestamp without time zone"`
}
