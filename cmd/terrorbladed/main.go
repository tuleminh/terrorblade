package main

import (
	"fmt"
	"log"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"

	"terrorblade/cmd/terrorbladed/internal/handlers"
	"terrorblade/kit/osenv"
	"terrorblade/kit/server"
)

func init() {
	fatalIfErrorOccurred(osenv.LoadFile("./conf/conf.default.env"))
}

func main() {
	m := InitializeMain()
	_ = m.Run()
}

// Sets definition.
var (
	AppSet = wire.NewSet(NewApp, NewMySQLConn, NewRoutingFn)
	KitSet = wire.NewSet(server.NewServer)
)

// NewApp returns a new instance of App.
func NewApp(server *server.Server) *App {
	return &App{
		Server: server,
	}
}

// App is main application.
type App struct {
	Server *server.Server
}

// Run starts application.
func (_this *App) Run() error {
	return _this.Server.Start()
}

// NewMySQLConn return *gorm.DB.
func NewMySQLConn() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v",
		osenv.String("DB_MYSQL_USERNAME"),
		osenv.String("DB_MYSQL_PASSWORD"),
		osenv.String("DB_MYSQL_HOST"),
		osenv.Int("DB_MYSQL_PORT"),
		osenv.String("DB_MYSQL_DATABASE"),
	))
	fatalIfErrorOccurred(err)
	return db
}

// NewRoutingFn returns routers.
func NewRoutingFn(h *handlers.Handlers) server.RoutingFn {
	return func(router *echo.Echo) {
		v1 := router.Group("/terrorblade/v1")
		v1.GET("/health-check", h.HealthCheck.Check())

		users := v1.Group("/users")
		{
			users.POST("/", h.User.CreateUser())
			users.GET("/:id", h.User.GetUser())
		}

		router.HTTPErrorHandler = func(err error, c echo.Context) {

		}
	}
}

func fatalIfErrorOccurred(err error) {
	if err != nil {
		log.Fatalf("An error occurred: %v", err)
	}
}
