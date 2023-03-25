package model

type AllCourseResponse struct {
	CourseCode    string   `json:"course_code"`
	Prerequisites []string `json:"prerequisites"`
}
