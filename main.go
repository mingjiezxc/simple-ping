package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	_ "github.com/mingjiezxc/simple-ping"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis没密码，没有设置，则留空
		DB:       0,                // 使用默认数据库
	})
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// list agent, agent info

// agent reset api
// agent add job
// agent del job
// agent post job
// job add ipGroup
// job del ipGroup

// add ipGroup
// ipGroup add ip
// ipGroup del ip
