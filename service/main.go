package main

import (
	"log"
	"service/config"
	"service/mysql"
	"service/router"
)

//func main() {
//	//r := gin.Default()
//	//r.POST("/course", handler.Get_course)
//	//r.Run()
//	err := crawler.Log_in("20S051030", "liguojian")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println("---------------------------------------------------------")
//	crawler.CrawlerCourse()
//}

//func helloHandler(c *gin.Context) {
//	claims := jwt.ExtractClaims(c)
//	user, _ := c.Get(identityKey)
//	c.JSON(200, gin.H{
//		"userID":   claims[identityKey],
//		"userName": user.(*User).UserName,
//		"text":     "Hello World.",
//	})
//}

func main() {
	conf := config.GetConfig()
	port := conf.Port
	mysql.InitDB()
	if err := router.Router().Run(":"+port); err != nil {
		log.Fatal(err)
	}
}
