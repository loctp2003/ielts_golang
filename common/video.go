package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Video struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url"`
	CloudName string `json:"cloud_name,omitempty"  gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Video) TableName() string { return "videos" }
func (j *Video) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmashal JSONB value:", value))

	}
	var vid Video
	if err := json.Unmarshal(bytes, &vid); err != nil {
		return err
	}
	*j = vid
	return nil

}
func (j *Video) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)

}

type Videos []Video

func (j *Videos) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmashal JSONB value:", value))

	}
	var vid Videos
	if err := json.Unmarshal(bytes, &vid); err != nil {
		return err
	}
	*j = vid
	return nil

}
func (j *Videos) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)

}
