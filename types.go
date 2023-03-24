package main

type Course struct {
	CourseCode    string   `bson:"course_code"`
	CourseName    string   `bson:"course_name"`
	Prerequisites []string `bson:"prerequisites"`
	Reviews       []any    `bson:"reviews"`
}
