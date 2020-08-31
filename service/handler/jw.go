package handler

import (
	"github.com/gin-gonic/gin"
)

func Get_course(c *gin.Context) {
	var data struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(401, gin.H{"message": "错误"})
		return
	}
	c.JSON(200, gin.H{"data": data})
}
