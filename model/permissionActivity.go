package model

import "time"

type PermissionActivity struct {
	tableName    struct{}   `pg:"auth.permission_activity"`
	PermissionId int        `json:"permission_id" pg:"type:integer"`
	ActivityId   int        `json:"activity_id" pg:"type:integer"`
	Permission   Permission `pg:"rel:has-one"`
	Activity     Activity   `pg:"rel:has-one"`
	CreatedAt    time.Time  `json:"created_at" pg:"type:timestamp without time zone,default:now()"`
	UpdatedAt    time.Time  `json:"updated_at" pg:"type:timestamp without time zone"`
}
