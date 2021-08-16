package main

import "github.com/gin-gonic/gin"

type GroupStruct struct {
	Name string   `json:"name" example:"group1"`
	IP   []string `json:"ip" example:"192.168.1.1,2.168.2.1"`
}

// @Summary IPGroup List
// @Description Get IPGroup List
// @Tags IP Group
// @Accept json
// @Produce json
// @Success 200 {array} string "Success"
// @Failure 201 {string} string "redis connect err"
// @Router /ip/group [get]
func IPGroupList(c *gin.Context) {
	val, err := rdb.Get(ctx, "group-list").Result()
	if PrintErr(err) {
		c.JSON(201, err)
		return
	}
	c.JSON(200, val)
}

// @Summary IPGroup IP List
// @Description Get IPGroup IP List
// @Tags IP Group
// @Accept json
// @Produce json
// @Param groupName path string true "Group Name"
// @Success 200 {array} string "Success"
// @Failure 201 {string} string "redis connect err"
// @Router /ip/group/{groupName} [get]
func IPGroupInfo(c *gin.Context) {
	groupName := c.Param("groupName")
	val, err := rdb.SMembers(ctx, GroupNameKey+groupName).Result()
	if PrintErr(err) {
		c.JSON(201, err)
		return
	}
	c.JSON(200, val)
}

// @Summary Create IPGroup and Add IP to Group
// @Description Create IPGroup and Add IP to Group
// @Tags IP Group
// @Accept json
// @Produce json
// @Param Group body GroupStruct true "Group Struct"
// @Success 200 {string} string "Success"
// @Failure 201 {string} string "redis connect err"
// @Failure 202 {string} string "request data err"
// @Router /ip/group [put]
func IPGroupAddIP(c *gin.Context) {
	var groupData GroupStruct
	if err := c.ShouldBindJSON(&groupData); PrintErr(err) {
		c.JSON(202, err)
		return
	}

	err := rdb.SAdd(ctx, GroupNameKey+groupData.Name, groupData.IP).Err()
	if PrintErr(err) {
		c.JSON(201, err)
		return
	}
	c.String(200, "Update Done")
}

// @Summary Delete IP On Group
// @Description Delete IP On Group
// @Tags IP Group
// @Accept json
// @Produce json
// @Param Group body GroupStruct true "Group Struct"
// @Success 200 {string} string "Success"
// @Failure 201 {string} string "redis connect err"
// @Failure 202 {string} string "request data err"
// @Router /ip/group [delete]
func IPGroupDelIP(c *gin.Context) {
	var groupData GroupStruct
	if err := c.ShouldBindJSON(&groupData); PrintErr(err) {
		c.JSON(202, err)
		return
	}

	err := rdb.SRem(ctx, GroupNameKey+groupData.Name, groupData.IP).Err()
	if PrintErr(err) {
		c.JSON(201, err)
		return
	}
	c.String(200, "group delete ip done ")
}

// @Summary Delete IP Group
// @Description Delete IP Group
// @Tags IP Group
// @Accept json
// @Produce json
// @Param groupName path string true "Group Name"
// @Success 200 {string} string "Success"
// @Failure 201 {string} string "redis connect err"
// @Router /ip/group/{groupName} [delete]
func IPGroupDelete(c *gin.Context) {
	groupName := c.Param("groupName")
	err := rdb.Del(ctx, GroupNameKey+groupName).Err()
	if PrintErr(err) {
		c.JSON(201, err)
		return
	}
	c.JSON(200, "delete group done")
}
