package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/opensourceways/sync-agent/client"
	"github.com/opensourceways/sync-agent/utils"
)

const (
	platformGitee  = "gitee"
	platformGithub = "github"
)

type baseController struct{}

func (bc *baseController) doResponse(c *gin.Context, resp interface{}, err error) {
	if err == nil {
		utils.SuccessWithData(c, resp)
	} else {
		utils.FailedWithError(c, err)
	}
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
