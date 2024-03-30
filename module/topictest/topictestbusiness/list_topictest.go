package topictestbusiness

import (
	"context"
	"ielts/common"
	"ielts/module/topictest/topictestmodel"
)

type ListTopicTestStore interface {
	ListDataByCondition(ctx context.Context,
		condition map[string]interface{},
		paging *common.Paging,
		moreKeys ...string,
	) ([]topictestmodel.TopicTest, error)
}
type listTopicTestBiz struct {
	store ListTopicTestStore
}

func NewListTopicTestBiz(store ListTopicTestStore) *listTopicTestBiz {
	return &listTopicTestBiz{store: store}
}
func (biz *listTopicTestBiz) ListTopicTest(
	ctx context.Context,
	paging *common.Paging,
) ([]topictestmodel.TopicTest, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, paging)
	return result, err

}
