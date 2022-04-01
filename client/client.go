package client

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/v36/github"
	"github.com/opensourceways/community-robot-lib/giteeclient"
	"github.com/opensourceways/community-robot-lib/secret"
	"github.com/opensourceways/go-gitee/gitee"
	"golang.org/x/oauth2"

	"github.com/opensourceways/sync-agent/config"
	"github.com/opensourceways/sync-agent/models"
)

var (
	githubClient githubSyncClient
	giteeClient  giteeSyncClient
)

type Client interface {
	SyncIssue(issue models.Issue) (*models.SyncIssueResult, error)
	SyncComment(comment models.Comment) error
	SyncIssueState(issue models.IssueUpdate) error
}

type githubSyncClient struct {
	ctx context.Context
	cli *github.Client
}

func (sc *githubSyncClient) SyncIssue(issue models.Issue) (*models.SyncIssueResult, error) {
	req := &github.IssueRequest{
		Title: &issue.Title,
		Body:  &issue.Content,
	}

	ci, _, err := sc.cli.Issues.Create(sc.ctx, issue.Org, issue.Repo, req)
	if err != nil {
		return nil, err
	}

	return &models.SyncIssueResult{
		OrgRepo: models.OrgRepo{
			Org:  issue.Org,
			Repo: issue.Repo,
		},
		Number: strconv.Itoa(ci.GetNumber()),
		Link:   ci.GetHTMLURL(),
	}, nil
}

func (sc *githubSyncClient) SyncComment(comment models.Comment) error {
	number, err := strconv.Atoi(comment.Number)
	if err != nil {
		return err
	}

	req := &github.IssueComment{
		Body: &comment.Content,
	}
	_, _, err = sc.cli.Issues.CreateComment(sc.ctx, comment.Org, comment.Repo, number, req)

	return err
}

func (sc *githubSyncClient) SyncIssueState(issue models.IssueUpdate) error {
	number, err := strconv.Atoi(issue.Number)
	if err != nil {
		return fmt.Errorf("%s is invalid github issue number: %v", issue.Number, err)
	}

	req := github.IssueRequest{State: &issue.State}
	_, _, err = sc.cli.Issues.Edit(sc.ctx, issue.Org, issue.Repo, number, &req)

	return err
}

type giteeSyncClient struct {
	ctx context.Context
	cli giteeclient.Client
}

func (sc *giteeSyncClient) SyncIssue(issue models.Issue) (*models.SyncIssueResult, error) {
	iss, err := sc.cli.CreateIssue(issue.Org, issue.Repo, issue.Title, issue.Content)
	if err != nil {
		return nil, err
	}

	return &models.SyncIssueResult{
		OrgRepo: models.OrgRepo{
			Org:  issue.Org,
			Repo: issue.Repo,
		},
		Number: iss.Number,
		Link:   iss.HtmlUrl,
	}, nil
}

func (sc *giteeSyncClient) SyncComment(comment models.Comment) error {
	return sc.cli.CreateIssueComment(comment.Org, comment.Repo, comment.Number, comment.Content)
}

func (sc *giteeSyncClient) SyncIssueState(issue models.IssueUpdate) error {
	param := gitee.IssueUpdateParam{
		Repo:  issue.Repo,
		State: issue.State,
	}
	_, err := sc.cli.UpdateIssue(issue.Org, issue.Number, param)

	return err
}

func Init() error {
	cfg := config.Config()
	giteeTP := cfg.GetGiteeTokenPath()
	githubTP := cfg.GetGithubTokenPath()

	agent := new(secret.Agent)
	if err := agent.Start([]string{giteeTP, githubTP}); err != nil {
		return err
	}

	defer agent.Stop()

	githubClient = githubSyncClient{
		ctx: context.Background(),
		cli: createGithubCli(agent.GetTokenGenerator(githubTP)),
	}

	giteeClient = giteeSyncClient{
		ctx: context.Background(),
		cli: giteeclient.NewClient(agent.GetTokenGenerator(giteeTP)),
	}

	return nil
}

func createGithubCli(getToken func() []byte) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: string(getToken())})
	tc := oauth2.NewClient(context.Background(), ts)

	return github.NewClient(tc)
}

func GiteeClient() Client {
	return &giteeClient
}

func GithubClient() Client {
	return &githubClient
}
