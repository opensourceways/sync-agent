package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/sync-agent/models"
)

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
	b := models.Comment{}
	if err := c.ShouldBind(&b); err != nil {
		logrus.Error(err)
		sc.responseBadRequest(c)

		return
	}

	client, err := platformClient(c.Param("platform"))
	if err != nil {
		logrus.Error(err)
		sc.responseBadRequest(c)

		return
	}

	if err := client.SyncComment(b); err != nil {
		logrus.Error(err)
		sc.responseFailed(c, codeSyncFailed, msgSyncFailed)

		return
	}

	sc.responseSuccess(c)
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
	b := models.Issue{}
	if err := c.ShouldBind(&b); err != nil {
		logrus.Error(err)
		sc.responseBadRequest(c)

		return
	}

	client, err := platformClient(c.Param("platform"))
	if err != nil {
		logrus.Error(err)
		sc.responseBadRequest(c)

		return
	}

	result, err := client.SyncIssue(b)
	if err != nil {
		logrus.Error(err)
		sc.responseFailed(c, codeSyncFailed, msgSyncFailed)

		return
	}

	sc.responseSuccessWithData(c, result)
}
