package gincourse

import (
	"github.com/gin-gonic/gin"
	"ielts/common"
	"ielts/component"
	"ielts/module/coursebusiness"
	"ielts/module/coursestorage"
	"net/http"
)

func ListCourse(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		paging.Fulfill()
		store := coursestorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := coursebusiness.NewListCourseBiz(store)
		result, err := biz.ListCoure(c.Request.Context(), &paging)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))

	}
}
