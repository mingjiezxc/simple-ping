package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	_ "simple-ping/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis没密码，没有设置，则留空
		DB:       0,                // 使用默认数据库
	})
	ctx = context.Background()
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
	r.GET("/ping", Ping)
	r.GET("/agents", AgentList)
	r.GET("/agents/:agentName", AgentList)
	r.GET("/agents/:agentName/jobs", AgentJobList)
	r.GET("/agents/:agentName/jobs/:jobName", AgentJobInfo)
	r.POST("/agents/:agentName/jobs/:jobName", AgentJobAdd)
	r.DELETE("/agents/:agentName/jobs/:jobName", AgentJobDel)

	r.GET("/ip/groups", IpGroupList)
	r.GET("/ip/groups/:groupName", IpGroupInfo)
	r.PUT("/ip/groups/:groupName", IpGroupAddIP)
	r.POST("/ip/groups/:groupName", IpGroupDelIP)
	r.DELETE("/ip/groups/:groupName", IpGroupDelete)

	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// @Summary a ping api
// @Description ping
// @Accept  text/plain
// @Success 200 {string} string	"pong"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.String(200, "pong")
}

func AgentList(c *gin.Context) {
	val, err := rdb.SMembers(ctx, "agent-list").Result()
	if PrintErr(err) {
		c.JSON(400, err)
		return
	}
	c.JSON(200, val)
}

func AgentOnline(c *gin.Context) {
	agentName := c.Param("agentName")
	val, err := rdb.Get(ctx, "agent-online_"+agentName).Result()
	if PrintErr(err) {
		c.JSON(400, err)
		return
	}
	c.JSON(200, val)
}

func AgentJobList(c *gin.Context) {
	agentName := c.Param("agentName")
	val, err := rdb.Get(ctx, "agent-jobs_"+agentName).Result()
	if PrintErr(err) {
		c.JSON(400, err)
		return
	}
	c.JSON(200, val)
}

func AgentJobInfo(c *gin.Context) {
	agentName := c.Param("agentName")
	jobName := c.Param("jobName")
	val, err := rdb.Get(ctx, "agent_"+agentName+"_"+jobName).Result()
	if PrintErr(err) {
		c.JSON(400, err)
		return
	}
	c.JSON(200, val)
}

func AgentJobAdd(c *gin.Context) {
	agentName := c.Param("agentName")
	jobName := c.Param("jobName")

	var jobData JobStruct
	if err := c.ShouldBindJSON(&jobData); err != nil {
		if PrintErr(err) {
			c.JSON(400, err)
			return
		}
	}

	jsobJson, err := json.Marshal(jobData)
	if PrintErr(err) {
		c.JSON(401, err)
		return
	}

	// add to agent jobs list
	_, err = rdb.SAdd(ctx, "agent-jobs_"+agentName, jobName).Result()
	if PrintErr(err) {
		c.JSON(402, err)
		return
	}

	// update or add jobs
	val, err := rdb.Set(ctx, "agent_"+agentName+"_"+jobName, jsobJson, 0).Result()
	if PrintErr(err) {
		c.JSON(403, err)
		return
	}
	c.JSON(200, val)
}

func AgentJobDel(c *gin.Context) {
	agentName := c.Param("agentName")
	jobName := c.Param("jobName")

	// remove job on agent jobs
	_, err := rdb.SRem(ctx, "agent-jobs_"+agentName, jobName).Result()
	if PrintErr(err) {
		c.JSON(400, err)
		return
	}
	// del job key
	val, err := rdb.Del(ctx, "agent_"+agentName+"_"+jobName).Result()
	if PrintErr(err) {
		c.JSON(401, err)
		return
	}
	c.JSON(200, val)
}

func IpGroupList(c *gin.Context) {
	val, err := rdb.Get(ctx, "groups-list").Result()
	if PrintErr(err) {
		c.JSON(400, err)
		return
	}
	c.JSON(200, val)
}

func IpGroupInfo(c *gin.Context) {
	groupName := c.Param("groupName")
	val, err := rdb.SMembers(ctx, "group_"+groupName).Result()
	if PrintErr(err) {
		c.JSON(400, err)
		return
	}
	c.JSON(200, val)
}

func IpGroupAddIP(c *gin.Context) {
	groupName := c.Param("groupName")
	val, err := rdb.SAdd(ctx, "group_"+groupName).Result()
	if PrintErr(err) {
		c.JSON(400, err)
		return
	}
	c.JSON(200, val)
}

func IpGroupDelIP(c *gin.Context) {
	groupName := c.Param("groupName")
	val, err := rdb.SRem(ctx, "group_"+groupName).Result()
	if PrintErr(err) {
		c.JSON(400, err)
		return
	}
	c.JSON(200, val)
}

func IpGroupDelete(c *gin.Context) {
	groupName := c.Param("groupName")
	val, err := rdb.Del(ctx, "group_"+groupName).Result()
	if PrintErr(err) {
		c.JSON(400, err)
		return
	}
	c.JSON(200, val)
}

type JobStruct struct {
	SPEC  string
	Name  string
	Group []string
}

func PrintErr(err error) bool {
	if err != nil {
		log.Println(err.Error())
		return true
	}
	return false
}
