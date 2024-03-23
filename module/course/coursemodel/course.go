package coursemodel

import (
	"errors"
	"ielts/common"
	"strings"
	"time"
)

const EntityName = "course"

type Course struct {
	Id              string        `json:"courseId" gorm:"column:courseId;"`
	status          int           `json:"status" gorm:"column:status;"`
	courseName      string        `json:"courseName" gorm:"column:courseName;"`
	description     string        `json:"description" gorm:"column:description;"`
	image           *common.Image `json:"image" gorm:"column:image;"`
	cost            int           `json:"cost" gorm:"column:cost;"`
	enrrollmentDate *time.Time    `json:"enrrollmentDate" gorm:"column:enrrollmentDate;"`
	trailer         *common.Video `json:"trailer" gorm:"column:trailer;"`
}

func (Course) TableName() string { return "course" }

type CourseUpdate struct {
	courseName      string        `json:"courseName" gorm:"column:courseName;"`
	description     string        `json:"description" gorm:"column:description;"`
	image           *common.Image `json:"image" gorm:"column:image;"`
	cost            int           `json:"cost" gorm:"column:cost;"`
	enrrollmentDate *time.Time    `json:"enrrollmentDate" gorm:"column:enrrollmentDate;"`
	trailer         *common.Video `json:"trailer" gorm:"column:trailer;"`
}
type CourseCreate struct {
	Id              string        `json:"courseId" gorm:"column:courseId;"`
	courseName      string        `json:"courseName" gorm:"column:courseName;"`
	description     string        `json:"description" gorm:"column:description;"`
	image           *common.Image `json:"image" gorm:"column:image;"`
	cost            int           `json:"cost" gorm:"column:cost;"`
	enrrollmentDate *time.Time    `json:"enrrollmentDate" gorm:"column:enrrollmentDate;"`
	trailer         *common.Video `json:"trailer" gorm:"column:trailer;"`
}

func (CourseUpdate) TableName() string {
	return Course{}.TableName()
}
func (CourseCreate) TableName() string {
	return Course{}.TableName()
}
func (cour *CourseCreate) Validate() error {
	cour.courseName = strings.TrimSpace(cour.courseName)
	if len(cour.courseName) == 0 {
		return errors.New("Course can be blank")
	}
	return nil
}
