package router

import (
	"fmt"
	//"github.com/gin-contrib/cors"
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"time"
)

type SettingRestApi struct {
	Environment string `json:"environment"`
	Host        string `json:"host"`
	Port        string `json:"port"`
}

const MAX_FILE_SIZE    = 8

func Load(setting SettingRestApi, routers ...func(group *gin.RouterGroup)) {
	listenAddress := fmt.Sprintf(":%s", setting.Port)
	router := gin.New()

	// setting router
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
	//	AllowHeaders:     []string{"Origin", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"},
	//	AllowCredentials: false,
	//	ExposeHeaders:    []string{"Content-Length"},
	//	MaxAge:           12 * time.Hour,
	//}))
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = int64(MAX_FILE_SIZE) << 20

	apiRouter := router.Group("/api/v1")

	for _, r := range routers {
		r(apiRouter)
	}

	//run
	router.Run(listenAddress)
}
