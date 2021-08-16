package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type IPStatus struct {
	Agent        string
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

// @Summary List Agnet Mon IP Status
// @Description  List Agnet Mon IP Status
// @Tags IP Group
// @Accept json
// @Produce json
// @Success 200 {array} string "Success"
// @Failure 201 {string} string "redis connect err"
// @Router /ip/status [get]
func IPStatusInfo(c *gin.Context) {
	agents, err := rdb.SMembers(ctx, AgentListKey).Result()
	if PrintErr(err) {
		c.JSON(201, err)
		return
	}

	var allIPStatus []IPStatus

	for _, agent := range agents {
		val, err := rdb.Get(ctx, AgentAllIPStatusKey+agent).Result()
		PrintErr(err)
		ipMap := make(map[string]IPStatus)
		err = json.Unmarshal([]byte(val), &ipMap)
		PrintErr(err)

		for _, status := range ipMap {
			allIPStatus = append(allIPStatus, status)
		}

	}

	c.JSON(200, allIPStatus)
}
