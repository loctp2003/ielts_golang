package gincourse

import (
	"github.com/gin-gonic/gin"
	"ielts/common"
	"ielts/component"
	"ielts/module/course/coursebusiness"
	"ielts/module/course/coursemodel"
	"ielts/module/course/coursestorage"
	"net/http"
)

func CreateCourse(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data coursemodel.CourseCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := coursestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := coursebusiness.NewCreateCourse(store)
		if err := biz.CreateCoure(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		data.GenUID(common.DbTypeCourse)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
