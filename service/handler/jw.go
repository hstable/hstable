package handler

import (
	"encoding/json"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"service/crawler"
	"service/model"
	"service/mysql"
	"time"
)

var IdentityKey = "account"

func PutCourseByJW(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	account := claims[IdentityKey]
	var data struct {
		Password string `json:"password"`
	}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(401, gin.H{"message": "struct Error!!!"})
		return
	}
	course_data, err := crawler.Log_in(account.(string), data.Password, true)
	if err != nil {
		fmt.Println(err)
		c.JSON(401, gin.H{"message": "login Error!!!"})
		return
	}
	c.JSON(200, gin.H{
		"course":      course_data,
		"latest_time": time.Now(),
	})
}

func GetCourseByDB(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	account := claims[IdentityKey]
	course_data, err := mysql.SelectByXh(account.(string))
	if err != nil {
		c.JSON(400, err)
		return
	}
	var stu_course model.Course
	json.Unmarshal([]byte(course_data.Course), &stu_course)
	c.JSON(200, gin.H{"course": model.StudentCourseResult{
		Id:             course_data.Id,
		Student_number: course_data.Student_number,
		Course:         stu_course,
		Last_sync:      course_data.Last_sync,
	}})
}
