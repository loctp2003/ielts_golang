package topictestmodel

import (
	"errors"
	"ielts/common"
	"strings"
)

const EntityName = "topictest"

type TopicTest struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"topicName" gorm:"column:topicName;"`
	Des             string        `json:"description" gorm:"column:description;"`
	Image           *common.Image `json:"image" gorm:"column:image;"`
}

func (TopicTest) TableName() string { return "topictest" }

type TopicTestUpdate struct {
	Name  string        `json:"topicName" gorm:"column:topicName;"`
	Des   string        `json:"description" gorm:"column:description;"`
	Image *common.Image `json:"image" gorm:"column:image;"`
}
type TopicTestCreate struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"topicName" gorm:"column:topicName;"`
	Des             string        `json:"description" gorm:"column:description;"`
	Image           *common.Image `json:"image" gorm:"column:image;"`
}

func (TopicTestUpdate) TableName() string {
	return TopicTest{}.TableName()
}
func (TopicTestCreate) TableName() string {
	return TopicTest{}.TableName()
}
func (topictest *TopicTestCreate) Validate() error {
	topictest.Name = strings.TrimSpace(topictest.Name)
	if len(topictest.Name) == 0 {
		return errors.New("Topictest can be blank")
	}
	return nil
}
func (data *TopicTest) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeCourse)
}
