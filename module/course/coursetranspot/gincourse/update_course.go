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

func UpdateCourse(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data coursemodel.CourseUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := coursestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := coursebusiness.NewUpdateCourseBiz(store)
		if err := biz.UpdateCorese(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
