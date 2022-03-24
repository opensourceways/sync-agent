package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/sync-agent/models"
)

//NotFoundError response access resource does not exist
func NotFoundError(c *gin.Context) {
	c.JSON(http.StatusOK, models.BaseResp{Code: CodeNotFound, Msg: CodeMsg(CodeNotFound)})
}

//SuccessWithData response success response with data
func SuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, models.BaseResp{Code: CodeSuccess, Msg: CodeMsg(CodeSuccess), Data: data})
}

// FailedWithError handle frontend return based on error
func FailedWithError(c *gin.Context, err error) {
	code := CodeServerCommError
	msg := CodeMsg(code)

	if err != nil {
		ce := errors.Cause(err)
		if e, ok := ce.(*XError); ok {
			code = e.ErrCode()
			msg = e.ErrMsg()
		}
	}

	logrus.Error(err)

	c.JSON(http.StatusBadRequest, models.BaseResp{Code: code, Msg: msg})
}
