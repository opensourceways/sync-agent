package models


// BaseResp is Server returns data based model
type BaseResp struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg" example:"请求成功/失败"`
	Data interface{} `json:"data"`
}


