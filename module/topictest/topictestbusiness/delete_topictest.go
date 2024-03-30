package topictestbusiness

import (
	"context"
	"ielts/common"
	"ielts/module/topictest/topictestmodel"
)

type DeleteTopicTestStore interface {
	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*topictestmodel.TopicTest, error)
}
type deleteTopictestBiz struct {
	store DeleteTopicTestStore
}

func NewDeleteTopicTestBiz(store DeleteTopicTestStore) *deleteTopictestBiz {
	return &deleteTopictestBiz{store: store}
}
func (biz *deleteTopictestBiz) DeleteTopicTest(
	ctx context.Context,
	id int,
) error {
	oldData, err := biz.store.FindDataByCondition(ctx, make(map[string]interface{}))
	if err != nil {
		return common.ErrCannotGetEntity(topictestmodel.EntityName, err)
	}
	if oldData.Status == 0 {
		return common.ErrEntityDeleted(topictestmodel.EntityName, err)
	}
	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(topictestmodel.EntityName, err)

	}
	return err
}
