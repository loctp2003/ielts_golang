package gintopictest

import (
	"github.com/gin-gonic/gin"
	"ielts/common"
	"ielts/component"
	"ielts/module/topictest/topictestbusiness"
	"ielts/module/topictest/topicteststorage"
	"net/http"
)

func GetCourse(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		strore := topicteststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := topictestbusiness.NewGetTopicTestBiz(strore)
		data, err := biz.GetTopicTest(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
