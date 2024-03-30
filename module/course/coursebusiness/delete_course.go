package coursebusiness

import (
	"context"
	"ielts/common"
	"ielts/module/course/coursemodel"
)

type DeleteCourseStore interface {
	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*coursemodel.Course, error)
}
type deleteCourseBiz struct {
	store DeleteCourseStore
}

func NewDeleteCourseBiz(store DeleteCourseStore) *deleteCourseBiz {
	return &deleteCourseBiz{store: store}
}
func (biz *deleteCourseBiz) DeleteCourse(
	ctx context.Context,
	id int,
) error {
	oldData, err := biz.store.FindDataByCondition(ctx, make(map[string]interface{}))
	if err != nil {
		return common.ErrCannotGetEntity(coursemodel.EntityName, err)
	}
	if oldData.Status == 0 {
		return common.ErrEntityDeleted(coursemodel.EntityName, err)
	}
	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(coursemodel.EntityName, err)

	}
	return err
}
