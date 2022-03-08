package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/opensourceways/sync-agent/utils"
)

type baseController struct{}

func (bc *baseController) ResponseSuccess(c *gin.Context) {
	utils.SuccessWithData(c, nil)
}

func (bc *baseController) ResponseSuccessWithData(c *gin.Context, data interface{}) {
	utils.SuccessWithData(c, data)
}


