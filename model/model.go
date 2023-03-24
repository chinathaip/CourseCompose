package model

type Course struct {
	CourseCode    string   `bson:"course_code" json:"course_code"`
	CourseName    string   `bson:"course_name" json:"course_name"`
	Programs      []string `bson:"programs" json:"programs"`
	Prerequisites []string `bson:"prerequisites" json:"prerequisites"`
	Reviews       []any    `bson:"reviews" json:"reviews"`
}
