package coursestorage

import (
	"context"
	"ielts/common"
	"ielts/module/course/coursemodel"
)

func (s *sqlStore) Create(ctx context.Context, data *coursemodel.CourseCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
