package coursestorage

import (
	"context"
	"ielts/common"
	"ielts/module/course/coursemodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	condition map[string]interface{},
	paging *common.Paging,
	moreKeys ...string,
) ([]coursemodel.Course, error) {
	db := s.db
	var result []coursemodel.Course
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	db = db.Table(coursemodel.CourseCreate{}.TableName()).
		Where(condition).
		Where("status in (1)")
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil

}
