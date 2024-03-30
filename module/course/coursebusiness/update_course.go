package coursebusiness

import (
	"context"
	"ielts/common"
	"ielts/module/course/coursemodel"
)

type UpdateCourseStore interface {
	UpdateDataByCondition(
		ctx context.Context,
		id int,
		data *coursemodel.CourseUpdate,
	) error
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*coursemodel.Course, error)
}
type updateCourseBiz struct {
	store UpdateCourseStore
}

func NewUpdateCourseBiz(store UpdateCourseStore) *updateCourseBiz {
	return &updateCourseBiz{store: store}
}
func (biz *updateCourseBiz) UpdateCorese(ctx context.Context, id int, data *coursemodel.CourseUpdate) error {
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
