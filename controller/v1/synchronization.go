package v1

import "github.com/gin-gonic/gin"

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
// @Param	platform	path	string	true	"平台：gitee 或 github"
// @Success 200 {object}	models.BaseResp "同步成功"
// @Failure	404	{object}	models.BaseResp	"错误返回"
// @Failure	500	{object}	models.BaseResp	"错误返回"
// @Router /synchronization/{platform}/comment [post]
func (sc *SyncController) Comment(c *gin.Context) {
	sc.ResponseSuccess(c)
}

// @Tags Synchronization
// @Summary 同步 gitee 或 github 平台的 issue
// @Produce json
// @Accept json
// @Param	platform	path	string	true	"平台：gitee 或 github"
// @Success 200 {object}	models.BaseResp "同步成功"
// @Failure	404	{object}	models.BaseResp	"错误返回"
// @Failure	500	{object}	models.BaseResp	"错误返回"
// @Router /synchronization/{platform}/issue [post]
func (sc *SyncController) Issue(c *gin.Context) {
	sc.ResponseSuccess(c)
}
