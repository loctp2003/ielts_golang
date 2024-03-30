package coursebusiness

import (
	"context"
	"ielts/module/course/coursemodel"
)

type CreateCourseStore interface {
	Create(ctx context.Context, data *coursemodel.CourseCreate) error
}
type createCourseBiz struct {
	store CreateCourseStore
}

func NewCreateCourse(store CreateCourseStore) *createCourseBiz {
	return &createCourseBiz{store: store}
}
func (biz *createCourseBiz) CreateCoure(ctx context.Context, data *coursemodel.CourseCreate) error {

	err := biz.store.Create(ctx, data)
	return err
}
