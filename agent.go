package main

import (
	"github.com/gin-gonic/gin"
)

// @Summary Agent List
// @Description get regedit agent list
// @Tags Agent
// @Accept json
// @Produce json
// @Success 200 {array} string "Success"
// @Failure 500 {string} string "redis connect err"
// @Router /agent [get]
func AgentList(c *gin.Context) {
	val, err := rdb.SMembers(ctx, AgentListKey).Result()
	if PrintErr(err) {
		c.JSON(500, err)
		return
	}
	c.JSON(200, val)
}

type IPStatus struct {
	IP           string
	PTLL         int64
	AllowedLoss  int64
	SendCount    int64
	ReceiveCount int64
	InErrList    bool
	Ms           []int64
	MsAvg        int64
	Loss         string
	Lost         int64
	UpdateTime   int64
}

// @Summary Agent IP Status
// @Description get Agent IP Status
// @Tags Agent
// @Accept json
// @Produce json
// @Param agentName path string true "Agent Name"
// @Success 200 {object} map[string]main.IPStatus "Success"
// @Failure 500 {string} string "redis connect err"
// @Router /agent/{agentName}/status [get]
func AgentAllIPStatus(c *gin.Context) {
	agentName := c.Param("agentName")
	val, err := rdb.Get(ctx, AgentAllIPStatusKey+agentName).Result()
	if PrintErr(err) {
		c.JSON(500, err)
		return
	}
	c.JSON(200, val)
}

// @Summary Agent Ping Err IP List
// @Description get Agent Ping Err IP List
// @Tags Agent
// @Accept json
// @Produce json
// @Param agentName path string true "Agent Name"
// @Success 200 {array} string "Success"
// @Failure 500 {string} string "redis connect err"
// @Router /agent/{agentName}/err/ip [get]
func AgentErrIPList(c *gin.Context) {
	agentName := c.Param("agentName")
	val, err := rdb.SMembers(ctx, AgentErrListKey+agentName).Result()
	if PrintErr(err) {
		c.JSON(500, err)
		return
	}
	c.JSON(200, val)
}

// @Summary Agent Ping IP Success Last ms
// @Description get Agent Ping IP Success last ms
// @Tags Agent
// @Accept json
// @Produce json
// @Param agentName path string true "Agent Name"
// @Param ip path string true "IP"
// @Success 200 {integer} integer "Success"
// @Failure 500 {string} string "redis connect err"
// @Router /agent/{agentName}/ip/{ip}/lastms [get]
func AgentIPLastMs(c *gin.Context) {
	agentName := c.Param("agentName")
	ip := c.Param("ip")
	val, err := rdb.Get(ctx, AgentIPLastMsKey+agentName+"_"+ip).Result()
	if PrintErr(err) {
		c.JSON(500, err)
		return
	}
	c.JSON(200, val)
}

// @Summary Agent Online status
// @Description get Agent Online status
// @Tags Agent
// @Accept json
// @Produce json
// @Param agentName path string true "Agent Name"
// @Success 200 {string} string "online"
// @Failure 500 {string} string "offline or redis connect err"
// @Router /agent/{agentName}/online [get]
func AgentOnline(c *gin.Context) {
	agentName := c.Param("agentName")
	err := rdb.Get(ctx, AgentOnlineKey+agentName).Err()
	if PrintErr(err) {
		c.String(500, "offline")
		return
	}
	c.String(200, "online")
}

// @Summary Agent Job List
// @Description get Agent Job List
// @Tags Agent
// @Accept json
// @Produce json
// @Param agentName path string true "Agent Name"
// @Success 200 {array} string "Success"
// @Failure 500 {string} string "redis connect err"
// @Router /agent/{agentName}/job [get]
func AgentJobList(c *gin.Context) {
	agentName := c.Param("agentName")
	val, err := rdb.SMembers(ctx, AgentJobListKey+agentName).Result()
	if PrintErr(err) {
		c.JSON(500, err)
		return
	}
	c.JSON(200, val)
}

// @Summary Agent Add Job
// @Description  Add Job to Agent JobList
// @Tags Agent
// @Accept json
// @Produce json
// @Param agentName path string true "Agent Name"
// @Param jobName path string true "Job Name"
// @Success 200 {array} string "Success"
// @Failure 500 {string} string "redis connect err"
// @Router /agent/{agentName}/job/{jobName} [post]
func AgentJobAdd(c *gin.Context) {
	agentName := c.Param("agentName")
	jobName := c.Param("jobName")

	// add to agent jobs list
	err := rdb.SAdd(ctx, AgentJobListKey+agentName, jobName).Err()
	if PrintErr(err) {
		c.JSON(402, err)
		return
	}

	c.String(200, "Agent add job Done")
}

// @Summary Agent Del Job
// @Description delete job on Agent JobList
// @Tags Agent
// @Accept json
// @Produce json
// @Param agentName path string true "Agent Name"
// @Param jobName path string true "Job Name"
// @Success 200 {string} string "Success"
// @Failure 500 {string} string "redis connect err"
// @Router /agent/{agentName}/job/{jobName} [delete]
func AgentJobDel(c *gin.Context) {
	agentName := c.Param("agentName")
	jobName := c.Param("jobName")

	// remove job on agent jobs
	_, err := rdb.SRem(ctx, AgentJobListKey+agentName, jobName).Result()
	if PrintErr(err) {
		c.JSON(500, err)
		return
	}

	c.String(200, "agent del job done ")
}
