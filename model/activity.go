package model

import "time"

type Activity struct {
	tableName struct{}  `pg:"auth.activities"` //bảng users có schema auth
	Id        int       `json:"id" pg:"type:serial,pk"`
	Url       string    `json:"url" pg:"type:varchar(255)"`
	Method    string    `json:"method" pg:"type:varchar(8),notnull"`
	UrlRegex  string    `json:"url_regex" pg:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" pg:"type:timestamp without time zone,default:now()"`
	UpdatedAt time.Time `json:"updated_at" pg:"type:timestamp without time zone"`

	Permissions []Permission `json:"permissions" pg:"rel:has-many"`
}
