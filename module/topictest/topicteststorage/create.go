package topicteststorage

import (
	"context"
	"ielts/common"
	"ielts/module/topictest/topictestmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *topictestmodel.TopicTestCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
