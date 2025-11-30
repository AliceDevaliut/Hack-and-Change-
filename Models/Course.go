package Models

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	gorm.Model
	Name string
	Description string
	Unit []Unit `gorm:"foreignKey:CourseRefer"`
	Assignments []Assignment `gorm:"foreignKey:CourseRefer"`
	TutorRefer int
}

type Unit struct {
	gorm.Model
	Name string
	Description string
	CourseRefer int
	Topics []Topic `gorm:"foreignKey:UnitRefer"`
}

type Topic struct {
	gorm.Model
	Name string
	Description string
	Articles []Article `gorm:"foreignKey:TopicRefer"`
	UnitRefer int
}

type Article struct {
	gorm.Model
	Name string
	Description string
	Path string
	TopicRefer int
}

type Assignment struct {
	gorm.Model
	Name string
	Description string
	CouseRefer int
	StartDate time.Time
	DueDate time.Time
	Submissions []Submission `gorm:"foreignKey:AssignmentRefer"`
}

type Submission struct {
	gorm.Model
	Name string
	Grade uint
	IsGraded bool
	AssignmentRefer int
	StudentRefer int
}