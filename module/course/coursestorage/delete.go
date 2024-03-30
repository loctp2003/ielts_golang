package coursestorage

import (
	"context"
	"ielts/module/course/coursemodel"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context,
	id int,
) error {
	db := s.db
	if err := db.Table(coursemodel.Course{}.TableName()).
		Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		panic(err)
	}
	return nil
}
