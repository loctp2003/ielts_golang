package topicteststorage

import (
	"context"
	"ielts/module/topictest/topictestmodel"
)

func (s *sqlStore) UpdateDataByCondition(ctx context.Context, id int, data *topictestmodel.TopicTest) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
