package coursestorage

import (
	"context"
	"ielts/module/course/coursemodel"
)

func (s *sqlStore) UpdateDataByCondition(ctx context.Context, id int, data *coursemodel.CourseUpdate) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
