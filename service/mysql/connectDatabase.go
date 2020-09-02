package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//数据库配置
const (
	userName = "root"
	password = "123456"
	ip = "db"
	port = "3306"
	dbName = "hstable"

)
//Db数据库连接池
var DB *sql.DB

func init() {
	InitDB()
}

//注意方法名大写，就是public
func InitDB()  {

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	//path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8mb4"}, "")
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", "mysql"}, "")
	fmt.Println("path: " + path)
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, err := sql.Open("mysql", path)
	if err != nil {
		panic(err)
	}

	create_db := "CREATE DATABASE IF NOT EXISTS hstable;"
	DB.Exec(create_db)
	use_db := "USE hstable"
	DB.Exec(use_db)
	create_table := `Create Table If Not Exists hstable.student_course(
	id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
		student_number VARCHAR(50) NOT NULL,
		course_info TEXT NOT NULL
	)DEFAULT CHARSET='utf8mb4';`
	DB.Exec(create_table)

	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}
