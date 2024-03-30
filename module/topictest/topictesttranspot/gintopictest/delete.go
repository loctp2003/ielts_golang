package gintopictest

import (
	"github.com/gin-gonic/gin"
	"ielts/common"
	"ielts/component"
	"ielts/module/topictest/topictestbusiness"
	"ielts/module/topictest/topicteststorage"
	"net/http"
)

func DeleteCourse(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := topicteststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := topictestbusiness.NewDeleteTopicTestBiz(store)
		if err := biz.DeleteTopicTest(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
