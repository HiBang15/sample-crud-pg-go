package sample_crud_pg_go

import (
	"github.com/HiBang15/sample-crud-pg-go.git/config"
	"github.com/HiBang15/sample-crud-pg-go.git/controller"
	"github.com/HiBang15/sample-crud-pg-go.git/router"
	"github.com/HiBang15/sample-crud-pg-go.git/router/public"
)

func main() {
	db := config.ConnectDataBase()  //Kết nối database
	defer db.Close()    //Đóng database trước kết thúc chương trình

	controller.DB = db

	settingRestApi := router.SettingRestApi{
		Environment: "debug",
		Host:        "127.0.0.1",
		Port:        "8080",
	}
	router.Load(settingRestApi, public.SetRouter)
}