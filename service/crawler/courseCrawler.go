package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"service/model"
	"service/mysql"
	"strings"
	//"github.com/go-sql-driver/mysql"
)

var JW_URL = "https://sso.hitsz.edu.cn:7002/cas/login?service=http://jw.hitsz.edu.cn/cas"
var Course_URL = "http://jw.hitsz.edu.cn/Xsxk/queryYxkc"
var JW_Mirror = "https://mr.hstable.cn:7002"
var Course_Mirror = "http://mr.hstable.cn:7003"

func get_lt(client *http.Client) (string, error) {
	var lt = ""
	resp, err := client.Get(JW_URL)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("..............")
		log.Println(err)
		return "error", err
	}
	//fmt.Print("lalala: \n" + string(body))
	//	提取校验码
	//fmt.Println(string(body))
	template := regexp.MustCompile(`<input.*?type="hidden".*?value="(.*?)".*?/>`)
	lt = template.FindStringSubmatch(string(body))[1]
	return lt, nil
}

func Log_in(account string, password string, forceUpdate bool) (model.Course, error) {
	var client = new(http.Client)
	client.Jar, _ = cookiejar.New(nil)
	// set mirror trickily
	client.Transport = &http.Transport{Dial: func(network string, addr string) (net.Conn, error) {
		_, port, err := net.SplitHostPort(addr)
		if err != nil {
			return nil, err
		}
		var target string
		// tricky
		if port == "7002" {
			target = JW_Mirror
		} else {
			target = Course_Mirror
		}
		u, err := url.Parse(target)
		if err != nil {
			panic(err)
		}
		ip := u.Hostname()
		port = u.Port()
		if port == "" {
			if u.Scheme == "https" {
				port = "443"
			} else {
				port = "80"
			}
		}
		if net.ParseIP(ip) == nil {
			ips, err := net.LookupHost(ip)
			if err != nil {
				return nil, err
			}
			ip = ips[0]
		}
		return net.Dial(network, net.JoinHostPort(ip, port))
	}}
	//params := model.PostParams{
	//	Username:   account,
	//	Password:   password,
	//	RememberMe: "off",
	//	Lt:         get_lt(client),
	//	Execution:  "e1s1",
	//	EventID:    "submit",
	//	VcUsername: "",
	//	VcPassword: "",
	//}
	lt, err := get_lt(client)
	if err != nil {
		return model.Course{}, err
	}
	params := url.Values{
		"username":    {account},
		"password":    {password},
		"rememberMe":  {"off"},
		"lt":          {lt},
		"execution":   {"e1s1"},
		"_eventId":    {"submit"},
		"vc_username": {""},
		"vc_password": {""},
	}
	//	发送请求
	resp, err := client.Post(JW_URL, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	body_str := string(body)
	//log.Println(body_str)
	if strings.Contains(body_str, "统一身份认证系统") {
		return model.Course{}, errors.New("login error")
	}
	defer func() {
		// logout
		res, err := client.Get("http://jw.hitsz.edu.cn/logout")
		if err == nil {
			res.Body.Close()
		}
	}()
	//fmt.Println(string(body_str))
	course_data := crawlerCourse(client)
	err = storeData(account, course_data, forceUpdate)
	if err != nil {
		return model.Course{}, err
	}
	return course_data, nil
}
func construct_params(params_json string) string {
	params_str := strings.Replace(params_json, "\"", "", -1)
	params_str = strings.Replace(params_str, "{", "", -1)
	params_str = strings.Replace(params_str, "}", "", -1)
	params_str = strings.Replace(params_str, ",", "&", -1)
	params_str = strings.Replace(params_str, ":", "=", -1)
	//fmt.Println("params_str: \n" + params_str)
	return params_str
}

func crawlerCourse(client *http.Client) model.Course {
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
	body_str := string(body)
	//log.Println(body_str)
	var course_data model.Course
	json.Unmarshal([]byte(body_str), &course_data)
	//fmt.Println(len(course_data.YxkcList))
	//fmt.Println(course_data.YxkcList[0])
	return course_data
}

func storeData(account string, course_data model.Course, force bool) error {
	// 存储到数据库
	course_info, _ := json.Marshal(course_data)
	v, err := mysql.SelectByXh(account)
	if err != nil {
		//fmt.Println(account)
		//fmt.Println(string(course_info))
		if course_data.YxkcList == nil {
			return errors.New("pulling course table failed")
		}
		mysql.Insert(account, string(course_info))
	} else if strings.TrimSpace(v.Course) == "" ||
		strings.TrimSpace(v.Course) == `{"yxkcList":null}` ||
		force {
		if course_data.YxkcList == nil {
			return errors.New("pulling course table failed")
		}
		mysql.UpdateByXh(account, string(course_info))
	}
	return nil
}
