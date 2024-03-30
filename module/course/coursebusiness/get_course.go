package coursebusiness

import (
	"context"
	"ielts/common"
	"ielts/module/course/coursemodel"
)

type GetCourseStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		morekeys ...string,
	) (*coursemodel.Course, error)
}
type getCourseBiz struct {
	store GetCourseStore
}

func NewGetCourseBiz(store GetCourseStore) *getCourseBiz {
	return &getCourseBiz{store: store}
}
func (biz *getCourseBiz) GetCourse(ctx context.Context, id int) (*coursemodel.Course, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{})
	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(coursemodel.EntityName, err)
		}
	}
	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(coursemodel.EntityName, err)
	}
	return data, err
}
