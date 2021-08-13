package main

import (
	"github.com/gin-gonic/gin"
)

func AgentList(c *gin.Context) {
	val, err := rdb.SMembers(ctx, AgentListKey).Result()
	if PrintErr(err) {
		c.JSON(400, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: val})
}

func AgentJobs(c *gin.Context) {
	agentName := c.Param("agentName")
	val, err := rdb.SMembers(ctx, AgentJobListKey+agentName).Result()
	if PrintErr(err) {
		c.JSON(400, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: val})
}

func AgentAllIPStatus(c *gin.Context) {
	agentName := c.Param("agentName")
	val, err := rdb.Get(ctx, AgentAllIPStatusKey+agentName).Result()
	if PrintErr(err) {
		c.JSON(400, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: val})
}

func AgentErrIPList(c *gin.Context) {
	agentName := c.Param("agentName")
	val, err := rdb.SMembers(ctx, AgentErrListKey+agentName).Result()
	if PrintErr(err) {
		c.JSON(400, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: val})
}

func AgentIPLastMs(c *gin.Context) {
	agentName := c.Param("agentName")
	ip := c.Param("ip")
	val, err := rdb.Get(ctx, AgentIPLastMsKey+agentName+"_"+ip).Result()
	if PrintErr(err) {
		c.JSON(400, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: val})
}

func AgentOnline(c *gin.Context) {
	agentName := c.Param("agentName")
	val, err := rdb.Get(ctx, AgentOnlineKey+agentName).Result()
	if PrintErr(err) {
		c.JSON(400, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: val})
}

func AgentJobList(c *gin.Context) {
	agentName := c.Param("agentName")
	val, err := rdb.SMembers(ctx, AgentJobListKey+agentName).Result()
	if PrintErr(err) {
		c.JSON(400, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: val})
}

func AgentJobAdd(c *gin.Context) {
	agentName := c.Param("agentName")
	jobName := c.Param("jobName")

	// add to agent jobs list
	_, err := rdb.SAdd(ctx, AgentJobListKey+agentName, jobName).Result()
	if PrintErr(err) {
		c.JSON(402, BaseReturn{Status: false, Data: err})
		return
	}

	c.JSON(200, BaseReturn{Status: true, Data: "Agent add job Done"})
}

func AgentJobDel(c *gin.Context) {
	agentName := c.Param("agentName")
	jobName := c.Param("jobName")

	// remove job on agent jobs
	_, err := rdb.SRem(ctx, AgentJobListKey+agentName, jobName).Result()
	if PrintErr(err) {
		c.JSON(400, BaseReturn{Status: false, Data: err})
		return
	}

	c.JSON(200, BaseReturn{Status: true, Data: "agent del job done "})
}
