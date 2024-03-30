package gincourse

import (
	"github.com/gin-gonic/gin"
	"ielts/common"
	"ielts/component"
	"ielts/module/course/coursebusiness"
	"ielts/module/course/coursestorage"
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
		result, err := biz.ListCourse(c.Request.Context(), &paging)
		if err != nil {
			panic(err)
		}
		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 { // shouldn't for client know logic how to get NextCursor
				paging.NextCursor = result[i].FakeId.String()
			}
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))

	}
}
