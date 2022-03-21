module github.com/opensourceways/sync-agent

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/google/go-github/v36 v36.0.0
	github.com/opensourceways/community-robot-lib v0.0.0-20220118064921-28924d0a1246
	github.com/sirupsen/logrus v1.8.1
	github.com/swaggo/gin-swagger v1.4.1
	github.com/swaggo/swag v1.7.9
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
)

replace github.com/opensourceways/community-robot-lib v0.0.0-20220118064921-28924d0a1246 => c:/goProject/community-robot-lib
