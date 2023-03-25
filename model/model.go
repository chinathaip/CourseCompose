package model

type Course struct {
	Code          string   `bson:"course_code" json:"course_code"`
	Name          string   `bson:"course_name" json:"course_name"`
	Programs      []string `bson:"programs" json:"programs"`
	Prerequisites []string `bson:"prerequisites" json:"prerequisites"`
	Reviews       []any    `bson:"reviews" json:"reviews"`
}
