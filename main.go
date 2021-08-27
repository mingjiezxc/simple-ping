package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v2"

	_ "simple-ping/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	rdb    *redis.Client
	config ConfigYaml
	ctx    = context.Background()
)

// redis key
var (
	// pls add AgentName
	AgentListKey        = "agent-list"
	AgentOnlineKey      = "agent-online_"
	AgentErrListKey     = "agent-err-list_"
	AgentJobListKey     = "agent-job-list_"
	AgentAllIPStatusKey = "agent-all-ip-status_"
	// pls add AgentName & ipAddr
	AgentIPLastMsKey = "agent-ip-last-ms_"
	GroupListKey     = "group-list"
	// pls add groupName
	GroupNameKey = "group_"
	JobListKey   = "job-list"
	// pls add jobName
	JobNameKey = "job_"
)

// @title Ping Agnet Manager Api
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService https://github.com/mingjiezxc

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func main() {

	// read config file
	configfile, err := ioutil.ReadFile("./config.yaml")
	if PrintErr(err) {
		os.Exit(1)
	}

	// yaml marshal config
	err = yaml.Unmarshal(configfile, &config)
	if PrintErr(err) {
		os.Exit(2)
	}

	// create redis client
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	r := gin.Default()
	// 配置模板
	r.LoadHTMLGlob("web/*")
	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	r.StaticFS("/static", http.Dir("./AdminLTE"))
	// r.StaticFS("/web", http.Dir("./web"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "agent.html", gin.H{})
	})

	r.GET("/job", func(c *gin.Context) {
		c.HTML(http.StatusOK, "job.html", gin.H{})
	})

	r.GET("/ip/group", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ipgroup.html", gin.H{})
	})

	r.GET("/ip/status", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ipstatus.html", gin.H{})
	})

	r.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.html", gin.H{})
	})

	v1Group := r.Group("/v1")
	v1Group.GET("/ping", Ping)
	v1Group.GET("/agent", AgentList)

	v1Group.GET("/agent/:agentName", AgentOnline)
	v1Group.GET("/agent/:agentName/status", AgentAllIPStatus)
	v1Group.GET("/agent/:agentName/err/ip", AgentErrIPList)
	v1Group.GET("/agent/:agentName/ip/:ip/lastms", AgentIPLastMs)

	v1Group.GET("/agent/:agentName/job", AgentJobList)
	v1Group.POST("/agent/:agentName/job/:jobName", AgentJobAdd)
	v1Group.DELETE("/agent/:agentName/job/:jobName", AgentJobDel)

	v1Group.GET("/job", JobList)
	v1Group.POST("/job", JobAdd)
	v1Group.DELETE("/job", JobDel)
	v1Group.GET("/job/:jobName", JobInfo)

	v1Group.GET("/ip/status", IPStatusInfo)

	v1Group.GET("/ip/group", IPGroupList)
	v1Group.PUT("/ip/group", IPGroupAddIP)
	v1Group.DELETE("/ip/group", IPGroupDelIP)
	v1Group.DELETE("/ip/group/:groupName", IPGroupDelete)
	v1Group.GET("/ip/group/:groupName", IPGroupInfo)

	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(config.ServerAddr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// @Summary a ping api
// @Description ping
// @Accept  text/plain
// @Success 200 {strng} string	"pong"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.String(200, "pong")
}

type BaseReturn struct {
	Status bool
	Data   interface{}
}

func PrintErr(err error) bool {
	if err != nil {
		log.Println(err.Error())
		return true
	}
	return false
}

type ConfigYaml struct {
	RedisAddr     string `yaml:"RedisAddr"`
	RedisPassword string `yaml:"RedisPassword"`
	RedisDB       int    `yaml:"RedisDB"`
	ServerAddr    string `yaml:"ServerAddr"`
}
