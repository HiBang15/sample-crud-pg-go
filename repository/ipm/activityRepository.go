package ipm

import (
	"github.com/HiBang15/sample-crud-pg-go.git/lib"
	"github.com/HiBang15/sample-crud-pg-go.git/model"
	"github.com/HiBang15/sample-crud-pg-go.git/repository"
	"github.com/go-pg/pg"
)

type ActivityRepository struct {
	Connector *pg.DB
}

const fileNameAR = "Activity Repository"

func NewActivityRepository(db *pg.DB) repository.IActivityRepository {
	return &ActivityRepository{Connector: db}
}

func (a *ActivityRepository) GetListActivity() (listActivities []*model.Activity, err error) {
	tx, err :=  a.Connector.Begin()

	err = a.Connector.Model(listActivities).Select()
	if err != nil {
		tx.Rollback()
		lib.WriteLogError(fileNameAR, "GetListActivity", err)
		return nil, err
	}
	tx.Commit()

	lib.WriteLogSuccess(fileNameAR, "GetListActivity")
	return listActivities, nil
}

func (a *ActivityRepository) GetActivityById(activityId int) (activity *model.Activity, err error) {
	err = a.Connector.Model(activity).Where("id = ?", activityId).Select()
	if err != nil {
		lib.WriteLogError(fileNameAR, "GetActivityById", err)
		return nil, err
	}

	lib.WriteLogSuccess(fileNameAR, "GetActivityById")
	return activity, nil
}

func (a *ActivityRepository) CreateActivity(activity *model.Activity) error {
	panic("implement me")
}

func (a *ActivityRepository) UpdateActivity(activity *model.Activity) error {
	panic("implement me")
}

func (a *ActivityRepository) DeleteActivity(activityId int) error {
	panic("implement me")
}
