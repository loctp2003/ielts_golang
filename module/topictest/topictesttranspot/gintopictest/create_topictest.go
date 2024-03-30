package gintopictest

import (
	"github.com/gin-gonic/gin"
	"ielts/common"
	"ielts/component"
	"ielts/module/topictest/topictestbusiness"
	"ielts/module/topictest/topictestmodel"
	"ielts/module/topictest/topicteststorage"
	"net/http"
)

func CreateCourse(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data topictestmodel.TopicTestCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := topicteststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := topictestbusiness.NewCreateTopicTest(store)
		if err := biz.CreateTopicTest(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		data.GenUID(common.DbTypeCourse)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
