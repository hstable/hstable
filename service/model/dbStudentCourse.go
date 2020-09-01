package model

type StudentCourse struct {
	Id int
	Student_number string
	Course string
}

type StudentCourseResult struct {
	Id int
	Student_number string
	Course Course
}
