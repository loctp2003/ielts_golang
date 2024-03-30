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

func UpdateCourse(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data topictestmodel.TopicTestUpdate
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
		store := topicteststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := topictestbusiness.NewUpdateTopicTestBiz(store)
		if err := biz.UpdateTopicTest(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
