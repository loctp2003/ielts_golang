package topicteststorage

import (
	"context"
	"ielts/module/topictest/topictestmodel"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context,
	id int,
) error {
	db := s.db
	if err := db.Table(topictestmodel.TopicTest{}.TableName()).
		Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		panic(err)
	}
	return nil
}
