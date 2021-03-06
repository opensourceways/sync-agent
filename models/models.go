package models

type OrgRepo struct {
	Org  string `json:"org" binding:"required"`
	Repo string `json:"repo" binding:"required"`
}

type Issue struct {
	OrgRepo
	Content string `json:"content" binding:"required"`
	Title   string `json:"title" binding:"required"`
}

type IssueUpdate struct {
	OrgRepo
	Number string `json:"number" binding:"required"`
	State  string `json:"state" binding:"required,oneof=open closed"`
}

type Comment struct {
	OrgRepo
	Number  string
	Content string
}

type SyncIssueResult struct {
	OrgRepo
	Number string
	Link   string
}
