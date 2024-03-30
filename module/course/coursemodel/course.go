package coursemodel

import (
	"errors"
	"ielts/common"
	"strings"
	"time"
)

const EntityName = "course"

type Course struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"courseName" gorm:"column:courseName;"`
	Des             string        `json:"description" gorm:"column:description;"`
	Image           *common.Image `json:"image" gorm:"column:image;"`
	Cost            int           `json:"cost" gorm:"column:cost;"`
	Date            *time.Time    `json:"enrrollmentDate" gorm:"column:enrrollmentDate;"`
	Trailer         *common.Video `json:"trailer" gorm:"column:trailer;"`
}

func (Course) TableName() string { return "course" }

type CourseUpdate struct {
	Name    string        `json:"courseName" gorm:"column:courseName;"`
	Des     string        `json:"description" gorm:"column:description;"`
	Image   *common.Image `json:"image" gorm:"column:image;"`
	Cost    int           `json:"cost" gorm:"column:cost;"`
	Date    *time.Time    `json:"enrrollmentDate" gorm:"column:enrrollmentDate;"`
	Trailer *common.Video `json:"trailer" gorm:"column:trailer;"`
}
type CourseCreate struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"courseName" gorm:"column:courseName;"`
	Des             string        `json:"description" gorm:"column:description;"`
	Image           *common.Image `json:"image" gorm:"column:image;"`
	Cost            int           `json:"cost" gorm:"column:cost;"`
	Date            *time.Time    `json:"enrrollmentDate" gorm:"column:enrrollmentDate;"`
	Trailer         *common.Video `json:"trailer" gorm:"column:trailer;"`
}

func (CourseUpdate) TableName() string {
	return Course{}.TableName()
}
func (CourseCreate) TableName() string {
	return Course{}.TableName()
}
func (cour *CourseCreate) Validate() error {
	cour.Name = strings.TrimSpace(cour.Name)
	if len(cour.Name) == 0 {
		return errors.New("Course can be blank")
	}
	return nil
}
func (data *Course) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeCourse)
}
