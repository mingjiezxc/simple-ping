package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type JobStruct struct {
	SPEC  string
	Name  string
	Group []string
}

func JobList(c *gin.Context) {
	val, err := rdb.SMembers(ctx, JobListKey).Result()
	if PrintErr(err) {
		c.JSON(500, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: val})
}

func JobInfo(c *gin.Context) {
	jobName := c.Param("jobName")
	val, err := rdb.Get(ctx, JobNameKey+jobName).Result()
	if PrintErr(err) {
		c.JSON(500, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: val})
}

func JobAdd(c *gin.Context) {

	var jobData JobStruct
	if err := c.ShouldBindJSON(&jobData); PrintErr(err) {
		c.JSON(500, BaseReturn{Status: false, Data: err})
		return
	}

	err := rdb.SAdd(ctx, JobListKey).Err()
	if PrintErr(err) {
		c.JSON(501, BaseReturn{Status: false, Data: err})
		return
	}

	jsonStr, _ := json.Marshal(jobData)

	err = rdb.Set(ctx, JobNameKey+jobData.Name, jsonStr, 0).Err()
	if PrintErr(err) {
		c.JSON(502, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: "add done"})
}

func JobDel(c *gin.Context) {
	var jobData JobStruct
	if err := c.ShouldBindJSON(&jobData); err != nil {
		if PrintErr(err) {
			c.JSON(500, BaseReturn{Status: false, Data: err})
			return
		}
	}

	err := rdb.SRem(ctx, JobListKey).Err()
	if PrintErr(err) {
		c.JSON(501, BaseReturn{Status: false, Data: err})
		return
	}

	err = rdb.Del(ctx, JobNameKey+jobData.Name).Err()
	if PrintErr(err) {
		c.JSON(502, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: ""})
}
