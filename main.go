package main

import (
	"flag"
	"os"

	"github.com/opensourceways/community-robot-lib/options"
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/sync-agent/core"
)

// @title Swagger sync-agent API
// @version 0.0.1
// @description plugin maintenance server api doc
// contact.name WeiZhi Xie
// contact.email 986740642@qq.com
// @securityDefinitions.apikey ApiKeyAuth
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @in header
// @name access-token
// @BasePath /v1
func main() {
	var opt options.ServiceOptions
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	opt.AddFlags(fs)
	if err := fs.Parse(os.Args[1:]); err != nil {
		logrus.WithError(err).Fatal("parse flag")
	}

	if err := opt.Validate(); err != nil {
		logrus.WithError(err).Fatal("invalid options")
	}

	core.Run(opt)
}
