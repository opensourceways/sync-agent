package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/opensourceways/sync-agent/models"
	"github.com/opensourceways/sync-agent/utils"
)

var paramError = utils.NewCodeError(utils.CodeBadeRequest)

type SyncController struct {
	baseController
}

func (sc *SyncController) Register(root *gin.RouterGroup) {
	syncRouter := root.Group("/synchronization")
	{
		syncRouter.POST("/:platform/comment", sc.Comment)
		syncRouter.POST("/:platform/issue", sc.Issue)
	}
}

// @Tags Synchronization
// @Summary 同步 gitee 或 github 平台的 comment
// @Produce json
// @Accept json
// @Param	platform	path	string			true	"平台：gitee 或 github"
// @Param	data		body	models.Comment	true	"需要同步的comment"
// @Success 200 {object}	models.BaseResp "同步成功"
// @Failure	404	{object}	models.BaseResp	"错误返回"
// @Failure	500	{object}	models.BaseResp	"错误返回"
// @Router /synchronization/{platform}/comment [post]
func (sc *SyncController) Comment(c *gin.Context) {
	err := func() error {
		b := models.Comment{}
		if err := c.ShouldBind(&b); err != nil {
			return errors.Wrap(paramError, err.Error())
		}

		client, err := platformClient(c.Param("platform"))
		if err != nil {
			return errors.Wrap(paramError, err.Error())
		}

		if err := client.SyncComment(b); err != nil {
			return errors.Wrapf(utils.NewCodeError(utils.CodeSyncIssueCommentFail), "sync comment error: %v", err)
		}

		return nil
	}()

	sc.doResponse(c, nil, err)
}

// @Tags Synchronization
// @Summary 同步 gitee 或 github 平台的 issue
// @Produce json
// @Accept json
// @Param	platform	path	string			true	"平台：gitee 或 github"
// @Param	data		body	models.Issue	true	"需要同步的issue"
// @Success 200 {object}	models.BaseResp{data=models.SyncIssueResult} "同步成功"
// @Failure	404	{object}	models.BaseResp	"错误返回"
// @Failure	500	{object}	models.BaseResp	"错误返回"
// @Router /synchronization/{platform}/issue [post]
func (sc *SyncController) Issue(c *gin.Context) {
	r, err := func() (*models.SyncIssueResult, error) {
		b := models.Issue{}
		if err := c.ShouldBind(&b); err != nil {
			return nil, errors.Wrap(paramError, err.Error())
		}

		client, err := platformClient(c.Param("platform"))
		if err != nil {
			return nil, errors.Wrap(paramError, err.Error())
		}

		result, err := client.SyncIssue(b)
		if err != nil {
			err = errors.Wrapf(utils.NewCodeError(utils.CodeSyncIssueFail), "sync issue fail: %v", err)
		}

		return result, err
	}()

	sc.doResponse(c, r, err)
}
