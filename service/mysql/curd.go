package mysql

import (
	"fmt"
	"service/model"
	"strconv"
)

func Insert(student_number string, course_info string) {
	insert_query := "insert into student_course (student_number, course_info) values(?,?)"
	rows, err := DB.Exec(insert_query, student_number, course_info)
	checkErr(err)
	num, err := rows.RowsAffected()
	checkErr(err)
	fmt.Println("effected: " + strconv.Itoa(int(num)))
}
func SelectByXh(student_number string) (model.StudentCourse, error) {
	var stu_course model.StudentCourse
	rows := DB.QueryRow("select  * from student_course where student_number = ?", student_number)
	err := rows.Scan(&stu_course.Id, &stu_course.Student_number, &stu_course.Course, &stu_course.Last_sync)
	if err != nil {
		return model.StudentCourse{}, err
	}
	//if stu_course.Id == 0 {
	//	fmt.Println("....................")
	//}
	//fmt.Println("select result: ", stu_course)
	return stu_course, nil
}
func DeleteByXh(student_number string) {
	rows, err := DB.Exec("delete from student_course where student_number=?", student_number)
	if err != nil {
		fmt.Println("delete error!", err)
	}
	num, _ := rows.RowsAffected()
	fmt.Println("effected: " + strconv.Itoa(int(num)))
}
func UpdateByXh(student_number string, course_info string) {
	rows, err := DB.Exec("update student_course set course_info=?, latest_update=NOW() where student_number=?", course_info, student_number)
	checkErr(err)
	num, _ := rows.RowsAffected()
	fmt.Println("affected number of rows: " + strconv.Itoa(int(num)))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
