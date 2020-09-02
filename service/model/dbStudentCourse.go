package model

import "time"

type StudentCourse struct {
	Id int
	Student_number string
	Course string
	Last_sync time.Time
}

type StudentCourseResult struct {
	Id int
	Student_number string
	Course Course
}
