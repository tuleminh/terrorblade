package server

import "github.com/labstack/echo/v4"

type RoutingFn func(router *echo.Echo)

// NewServer returns a new instance of Server.
func NewServer(routing RoutingFn) *Server {
	e := echo.New()
	e.HideBanner = true
	return &Server{
		router:  e,
		routing: routing,
	}
}

type Server struct {
	router  *echo.Echo
	routing RoutingFn
}

func (_this *Server) Start() error {
	_this.routing(_this.router)

	return _this.router.Start(":8080")
}
