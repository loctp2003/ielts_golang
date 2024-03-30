package topictestbusiness

import (
	"context"
	"ielts/module/topictest/topictestmodel"
)

type CreateTopictestStore interface {
	Create(ctx context.Context, data *topictestmodel.TopicTestCreate) error
}
type createTopicTestBiz struct {
	store CreateTopictestStore
}

func NewCreateTopicTest(store CreateTopictestStore) *createTopicTestBiz {
	return &createTopicTestBiz{store: store}
}
func (biz *createTopicTestBiz) CreateTopicTest(ctx context.Context, data *topictestmodel.TopicTestCreate) error {

	err := biz.store.Create(ctx, data)
	return err
}
