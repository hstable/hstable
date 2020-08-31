package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"service/model"
	"service/mysql"
	"strings"
	//"github.com/go-sql-driver/mysql"
)

var client http.Client

func init() {
	client.Jar, _ = cookiejar.New(nil)
}



var JW_URL = "https://sso.hitsz.edu.cn:7002/cas/login?service=http://jw.hitsz.edu.cn/casLogin"
var Course_URL = "http://jw.hitsz.edu.cn/Xsxk/queryYxkc"

func get_lt() string {
	resp, err := client.Get(JW_URL)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	//fmt.Print("lalala: \n" + string(body))
	//	提取校验码
	template := regexp.MustCompile(`<input.*?type="hidden".*?value="(.*?)".*?/>`)
	lt := template.FindStringSubmatch(string(body))[1]
	return lt
}

func Log_in(account string, password string) {
	params := model.PostParams{
		Username:   account,
		Password:   password,
		RememberMe: "on",
		Lt:         get_lt(),
		Execution:  "e1s1",
		EventID:    "submit",
		VcUsername: "",
		VcPassword: "",
	}
	params_json, err := json.Marshal(params)
	if err != nil {
		log.Println(err)
	}
	params_str := construct_params(string(params_json))
	//	发送请求
	resp, err := client.Post(JW_URL, "application/x-www-form-urlencoded", strings.NewReader(params_str))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(body))
}
func construct_params(params_json string) string {
	params_str := strings.Replace(params_json, "\"", "", -1)
	params_str = strings.Replace(params_str, "{", "", -1)
	params_str = strings.Replace(params_str, "}", "", -1)
	params_str = strings.Replace(params_str, ",", "&", -1)
	params_str = strings.Replace(params_str, ":", "=", -1)
	fmt.Println("params_str: \n" + params_str)
	return params_str
}

func CrawlerCourse() model.Course {
	params := model.Get_Post_Course()
	params_json, err := json.Marshal(params)
	if err != nil {
		log.Println(err)
	}
	params_str := construct_params(string(params_json))
	resp, err := client.Post(Course_URL, "application/x-www-form-urlencoded; charset=UTF-8", strings.NewReader(params_str))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var course_data model.Course
	json.Unmarshal([]byte(string(body)), &course_data)
	fmt.Println(len(course_data.YxkcList))
	fmt.Println(course_data.YxkcList[1])
	// 存储到数据库
	student_number := course_data.YxkcList[0].Xh
	course_info, err := json.Marshal(course_data)
	stu_course := mysql.SelectByXh(student_number)
	if stu_course.Id == 0 {
		mysql.Insert(student_number, string(course_info))
	}else {
		mysql.UpdateByXh(student_number, string(course_info))
	}
	return course_data
}