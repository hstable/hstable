package main

import (
	"fmt"
	"service/crawler"
)

func main() {
	//r := gin.Default()
	//r.POST("/course", handler.Get_course)
	//r.Run()
	crawler.Log_in("20S051030", "liguojian123")
	fmt.Println("---------------------------------------------------------")
	crawler.CrawlerCourse()
}
