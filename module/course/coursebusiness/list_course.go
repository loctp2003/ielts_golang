package coursebusiness

import (
	"context"
	"ielts/common"
	"ielts/module/course/coursemodel"
)

type ListCourseStore interface {
	ListDataByCondition(ctx context.Context,
		condition map[string]interface{},
		paging *common.Paging,
		moreKeys ...string,
	) ([]coursemodel.Course, error)
}
type listCourseBiz struct {
	store ListCourseStore
}

func NewListCourseBiz(store ListCourseStore) *listCourseBiz {
	return &listCourseBiz{store: store}
}
func (biz *listCourseBiz) ListCourse(
	ctx context.Context,
	paging *common.Paging,
) ([]coursemodel.Course, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, paging)
	return result, err

}
