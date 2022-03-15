package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opensourceways/sync-agent/models"
)

const (
	CodeSuccess     = 200
	CodeNotFound    = 404
	CodeBadeRequest = 400
)

const (
	MsgSuccess    = "请求成功"
	MsgNotFound   = "访问的资源不存在"
	MsgBadRequest = "请求参数错误"
)

//NotFoundError response access resource does not exist
func NotFoundError(c *gin.Context) {
	c.JSON(http.StatusOK, models.BaseResp{Code: CodeNotFound, Msg: MsgNotFound})
}

//SuccessWithData response success response with data
func SuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, models.BaseResp{Code: CodeSuccess, Msg: MsgSuccess, Data: data})
}

func BadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, models.BaseResp{Code: CodeBadeRequest, Msg: MsgBadRequest,})
}
