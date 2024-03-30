package gintopictest

import (
	"github.com/gin-gonic/gin"
	"ielts/common"
	"ielts/component"
	"ielts/module/topictest/topictestbusiness"
	"ielts/module/topictest/topicteststorage"
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
		store := topicteststorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := topictestbusiness.NewListTopicTestBiz(store)
		result, err := biz.ListTopicTest(c.Request.Context(), &paging)
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
