package config

import (
	"errors"
	"sync"

	libconfig "github.com/opensourceways/community-robot-lib/config"
)

var (
	once  sync.Once
	agent libconfig.ConfigAgent
)

type configuration struct {
	Env             string `json:"env"`
	GiteeTokenPath  string `json:"gitee_token_path",required:"true"`
	GithubTokenPath string `json:"github_token_path",required:"true"`
}

func (c *configuration) SetDefault() {
	if c.Env == "" {
		c.Env = "dev"
	}
}

func (c *configuration) Validate() error {
	if c.GiteeTokenPath == "" {
		return errors.New("the gitee_token_path configuration filed can't be empty")
	}

	if c.GithubTokenPath == "" {
		return errors.New("the github_token_path configuration filed can't be empty")
	}

	return nil
}

func (c *configuration) GetEnv() string {
	if c != nil {
		return c.Env
	}

	return ""
}

func (c *configuration) GetGiteeTokenPath() string {
	if c != nil {
		return c.GiteeTokenPath
	}

	return ""
}

func (c *configuration) GetGithubTokenPath() string {
	if c != nil {
		return c.GithubTokenPath
	}

	return ""
}

func newConfig() libconfig.Config {
	return &configuration{}
}

func Init(path string) (err error) {
	once.Do(func() {
		agent = libconfig.NewConfigAgent(newConfig)
		err = agent.Start(path)
	})

	return
}

func Release() {
	agent.Stop()
}

func Config() *configuration {
	_, c := agent.GetConfig()
	if cfg, ok := c.(*configuration); ok {
		return cfg
	}

	return nil
}
