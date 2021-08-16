package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type JobStruct struct {
	SPEC  string   `json:"spec" example:"*/10 * * * * *"`
	Name  string   `json:"name" example:"job1"`
	Group []string `json:"group" example:"group1,group2"`
}

// @Summary  Job List
// @Description get Job List
// @Tags Job
// @Accept json
// @Produce json
// @Success 200 {array} string "Success"
// @Failure 500 {string} string "redis connect err"
// @Router /job [get]
func JobList(c *gin.Context) {
	val, err := rdb.SMembers(ctx, JobListKey).Result()
	if PrintErr(err) {
		c.JSON(500, err)
		return
	}
	c.JSON(200, val)
}

// @Summary Get Job Info
// @Description Get Job Info
// @Tags Job
// @Accept json
// @Produce json
// @Success 200 {object} JobStruct "Success"
// @Failure 500 {string} string "redis connect err"
// @Router /job/{jobName} [get]
func JobInfo(c *gin.Context) {
	jobName := c.Param("jobName")
	val, err := rdb.Get(ctx, JobNameKey+jobName).Result()
	if PrintErr(err) {
		c.JSON(500, err)
		return
	}
	c.JSON(200, val)
}

// @Summary Add & Update Job
// @Description Add & Update Job
// @Tags Job
// @Accept json
// @Produce json
// @Success 200 {object} JobStruct "Success"
// @Failure 500 {string} string "request data err"
// @Failure 501 {string} string "redis connect err"
// @Failure 502 {string} string "redis connect err"
// @Router /job [post]
func JobAdd(c *gin.Context) {

	var jobData JobStruct
	if err := c.ShouldBindJSON(&jobData); PrintErr(err) {
		c.JSON(500, err)
		return
	}

	// add job to jobList
	err := rdb.SAdd(ctx, JobListKey).Err()
	if PrintErr(err) {
		c.JSON(501, BaseReturn{Status: false, Data: err})
		return
	}

	jsonStr, _ := json.Marshal(jobData)

	// add & update job
	err = rdb.Set(ctx, JobNameKey+jobData.Name, jsonStr, 0).Err()
	if PrintErr(err) {
		c.JSON(502, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: "add done"})
}

// @Summary Delete Job
// @Description Delete Job
// @Tags Job
// @Accept json
// @Produce json
// @Success 200 {object} JobStruct "Success"
// @Failure 500 {string} string "request data err"
// @Failure 501 {string} string "redis connect err"
// @Failure 502 {string} string "redis connect err"
// @Router /job [delete]
func JobDel(c *gin.Context) {
	var jobData JobStruct
	if err := c.ShouldBindJSON(&jobData); err != nil {
		if PrintErr(err) {
			c.JSON(500, err)
			return
		}
	}

	err := rdb.SRem(ctx, JobListKey, jobData.Name).Err()
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
