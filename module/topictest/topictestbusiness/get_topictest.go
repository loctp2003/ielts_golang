package topictestbusiness

import (
	"context"
	"ielts/common"
	"ielts/module/topictest/topictestmodel"
)

type GetTopicTestStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		morekeys ...string,
	) (*topictestmodel.TopicTest, error)
}
type getTopicTestBiz struct {
	store GetTopicTestStore
}

func NewGetTopicTestBiz(store GetTopicTestStore) *getTopicTestBiz {
	return &getTopicTestBiz{store: store}
}
func (biz *getTopicTestBiz) GetTopicTest(ctx context.Context, id int) (*topictestmodel.TopicTest, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{})
	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(topictestmodel.EntityName, err)
		}
	}
	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(topictestmodel.EntityName, err)
	}
	return data, err
}
