package repository

import "github.com/HiBang15/sample-crud-pg-go.git/model"

type IActivityRepository interface {
	GetListActivity() ([]*model.Activity, error)
	GetActivityById(int) (*model.Activity, error)
	CreateActivity(*model.Activity) error
	UpdateActivity(*model.Activity) error
	DeleteActivity(int) error
}
