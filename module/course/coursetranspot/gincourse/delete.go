package gincourse

import (
	"github.com/gin-gonic/gin"
	"ielts/common"
	"ielts/component"
	"ielts/module/course/coursebusiness"
	"ielts/module/course/coursestorage"
	"net/http"
)

func DeleteCourse(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := coursestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := coursebusiness.NewDeleteCourseBiz(store)
		if err := biz.DeleteCourse(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
