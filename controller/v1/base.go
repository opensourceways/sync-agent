package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/opensourceways/sync-agent/client"
	"github.com/opensourceways/sync-agent/models"
	"github.com/opensourceways/sync-agent/utils"
)

const (
	platformGitee  = "gitee"
	platformGithub = "github"
)

const (
	codeSyncFailed = 4001
	msgSyncFailed  = "同步失败"
)

type baseController struct{}

func (bc *baseController) responseSuccess(c *gin.Context) {
	utils.SuccessWithData(c, nil)
}

func (bc *baseController) responseFailed(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusBadRequest, models.BaseResp{Code: code, Msg: msg})
}

func (bc *baseController) responseSuccessWithData(c *gin.Context, data interface{}) {
	utils.SuccessWithData(c, data)
}

func (bc *baseController) responseBadRequest(c *gin.Context) {
	utils.BadRequest(c)
}

func platformClient(p string) (client.Client, error) {
	switch p {
	case platformGitee:
		return client.GiteeClient(), nil
	case platformGithub:
		return client.GithubClient(), nil
	}
	return nil, fmt.Errorf("does not support %s platform", p)
}
