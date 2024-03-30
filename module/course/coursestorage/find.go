package coursestorage

import (
	"context"
	"gorm.io/gorm"
	"ielts/common"
	"ielts/module/course/coursemodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*coursemodel.Course, error) {
	var result coursemodel.Course
	db := s.db
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	if err := db.Where(condition).
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
