package router

import "github.com/gin-gonic/gin"

type IRouter interface {
	Register(group *gin.RouterGroup)
}

type Options struct {
	routeGP      *gin.RouterGroup
	routeModules []IRouter
}

type Option func(options *Options)

// RouteModules set the routing module that needs to be registered
func RouteModules(rm ...IRouter) Option {
	return func(options *Options) {
		options.routeModules = rm
	}
}

// RouteGP setting router group
func RouteGP(gp *gin.RouterGroup) Option {
	return func(options *Options) {
		options.routeGP = gp
	}
}

// RouteRegister register routing and manage routing packets
type RouteRegister struct {
	Options
}

func (reg *RouteRegister) AddRouter(routes ...IRouter) *RouteRegister {
	reg.routeModules = append(reg.routeModules, routes...)

	return reg
}

func (reg *RouteRegister) Do() {
	for _, r := range reg.routeModules {
		r.Register(reg.routeGP)
	}
}

func NewRouteRegister(opts ...Option) *RouteRegister {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	if opt.routeGP == nil {
		opt.routeGP = gin.Default().Group("/")
	}

	return &RouteRegister{Options: opt}
}
