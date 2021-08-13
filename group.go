package main

import "github.com/gin-gonic/gin"

type GroupStruct struct {
	Name string
	IP   []string
}

func GroupList(c *gin.Context) {
	val, err := rdb.Get(ctx, "group-list").Result()
	if PrintErr(err) {
		c.JSON(400, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: val})
}

func GroupInfo(c *gin.Context) {
	groupName := c.Param("groupName")
	val, err := rdb.SMembers(ctx, GroupNameKey+groupName).Result()
	if PrintErr(err) {
		c.JSON(400, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: val})
}

func GroupAddIP(c *gin.Context) {
	var groupData GroupStruct
	if err := c.ShouldBindJSON(&groupData); PrintErr(err) {
		c.JSON(500, BaseReturn{Status: false, Data: err})
		return
	}

	val, err := rdb.SAdd(ctx, GroupNameKey+groupData.Name, groupData.IP).Result()
	if PrintErr(err) {
		c.JSON(400, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: val})
}

func GroupDelIP(c *gin.Context) {
	var groupData GroupStruct
	if err := c.ShouldBindJSON(&groupData); PrintErr(err) {
		c.JSON(500, BaseReturn{Status: false, Data: err})
		return
	}

	val, err := rdb.SRem(ctx, GroupNameKey+groupData.Name, groupData.IP).Result()
	if PrintErr(err) {
		c.JSON(400, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: val})
}

func GroupDelete(c *gin.Context) {
	groupName := c.Param("groupName")
	val, err := rdb.Del(ctx, GroupNameKey+groupName).Result()
	if PrintErr(err) {
		c.JSON(400, BaseReturn{Status: false, Data: err})
		return
	}
	c.JSON(200, BaseReturn{Status: true, Data: val})
}
