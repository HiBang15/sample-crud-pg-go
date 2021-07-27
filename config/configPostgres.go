package config

import (
	"github.com/HiBang15/sample-crud-pg-go.git/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
)

//var connector *pg.DB

func ConnectDataBasePostgres() (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		Addr:                  ":5432",
		User:                  "postgres",
		Password:              "secret",
		Database:              "postgres",
	})

	err := createSchema(db)
	if err != nil {
		log.Println("[Config][ConnectDatabase][Success] Connect database fail with err: " + err.Error())
		log.Fatal(err)
	}
	log.Println("[Config][ConnectDatabase][Success] Connect database success!")

	return db
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*model.User)(nil),
		(*model.Role)(nil),
		(*model.UserRole)(nil),
		(*model.Permission)(nil),
		(*model.RolePermission)(nil),
		(*model.Activity)(nil),
		(*model.PermissionActivity)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false,        //Nếu đặt Temp: true sẽ biến thành in-memory database
			IfNotExists: true,  //Nếu tồn tại bảng sẽ không tạo thêm
			Varchar: 255,
		})

		if err != nil {
			log.Println("[Config][CreateSchema][Error] create schema fail with err: " + err.Error())
			return err
		}
	}
	log.Println("[Config][CreateSchema][Success]create schema successful!")
	return nil
}
