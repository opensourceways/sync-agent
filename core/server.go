package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/opensourceways/community-robot-lib/interrupts"
	"github.com/opensourceways/community-robot-lib/logrusutil"
	"github.com/opensourceways/community-robot-lib/options"
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/sync-agent/client"
	"github.com/opensourceways/sync-agent/config"
)

func init() {
	logrusutil.ComponentInit("sync-agent")
}

func Run(opt options.ServiceOptions) {
	if err := config.Init(opt.ConfigFile); err != nil {
		logrus.WithError(err).Fatal("init config fail")
	}

	defer config.Release()

	cfg := config.Config()
	if cfg == nil {
		logrus.Fatal("can't get configuration")
		return
	}

	if cfg.GetEnv() == "release" {
		logrus.SetLevel(logrus.ErrorLevel)

	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if err := client.Init(); err != nil {
		logrus.WithError(err).Fatal("init client")
		return
	}

	defer interrupts.WaitForGracefulShutdown()

	logrus.Info(fmt.Printf(
		"welcome sync-agent api. the default doc address: http://localhost:%d/swagger/index.html\n",
		opt.Port,
	))

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", opt.Port),
		Handler:        initRouter(),
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	interrupts.ListenAndServe(s, opt.GracePeriod)
}
