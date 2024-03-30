package gincourse

import (
	"github.com/gin-gonic/gin"
	"ielts/common"
	"ielts/component"
	"ielts/module/course/coursebusiness"
	"ielts/module/course/coursestorage"
	"net/http"
)

func GetCourse(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		strore := coursestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := coursebusiness.NewGetCourseBiz(strore)
		data, err := biz.GetCourse(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
