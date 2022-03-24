package utils

const (
	CodeSuccess         uint32 = 200
	CodeNotFound        uint32 = 404
	CodeBadeRequest     uint32 = 400
	CodeServerCommError uint32 = 1001
)

const (
	CodeSyncIssueFail        uint32 = 2001
	CodeSyncIssueCommentFail uint32 = 2002
)

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[CodeSuccess] = "请求成功"
	message[CodeBadeRequest] = "请求参数错误"
	message[CodeNotFound] = "请求资源不存在"
	message[CodeServerCommError] = "服务器开小差了，请稍后再试"
	message[CodeSyncIssueFail] = "同步issue失败"
	message[CodeSyncIssueCommentFail] = "同步issue评论失败"
}

func CodeMsg(code uint32) string {
	if v, ok := message[code]; ok {
		return v
	}

	return "unknown error"
}

func IsMsgCode(code uint32) bool {
	_, ok := message[code]

	return ok
}
