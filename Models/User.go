package Models

import (
	"gorm.io/gorm"
)

type Tutor struct {
	gorm.Model
	Login string
	Name string
	Courses []Course `gorm:"foreignKey:TutorRefer"`
}

type Student struct {
	gorm.Model
	Login string
	Name string
	Assignments []Assignment `gorm:"foreignKey:StudentRefer"`
	
}