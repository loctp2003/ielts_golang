package topictestbusiness

import (
	"context"
	"ielts/common"
	"ielts/module/topictest/topictestmodel"
)

type UpdateTopicTestStore interface {
	UpdateDataByCondition(
		ctx context.Context,
		id int,
		data *topictestmodel.TopicTestUpdate,
	) error
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*topictestmodel.TopicTest, error)
}
type updateTopicTestBiz struct {
	store UpdateTopicTestStore
}

func NewUpdateTopicTestBiz(store UpdateTopicTestStore) *updateTopicTestBiz {
	return &updateTopicTestBiz{store: store}
}
func (biz *updateTopicTestBiz) UpdateTopicTest(ctx context.Context, id int, data *topictestmodel.TopicTestUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{})
	if err != nil {
		if err != common.RecordNotFound {
			return nil
		}
	}
	if oldData.Status == 0 {
		return nil
	}
	if err := biz.store.UpdateDataByCondition(
		ctx,
		id,
		data,
	); err != nil {
	}
	return err
}
