package public

import (
	"github.com/HiBang15/sample-crud-pg-go.git/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func SetRouter(router *gin.RouterGroup) {
	log.Print("Start init public router .....")

	user := router.Group("user")
	{
		user.GET("/all", controller.GetAllUser)
		user.GET("/detail/:id", controller.GetUserById)
		user.POST("/", controller.CreateUser)
		user.PUT("/:id", controller.GetAllUser)
		user.DELETE("/:id", controller.GetAllUser)
	}

	log.Print("Finish init public router ....")
}
