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
	err := osenv.LoadFile("./conf/conf.default.env")
	checkError(err)

	fmt.Println(osenv.String("DB_MYSQL_USERNAME"))
}

func main() {
	m := InitializeMain()
	_ = m.Run()
}

var (
	MainSet = wire.NewSet(NewApp, NewMySQLConnection, NewRoutingFn)
	KitSet  = wire.NewSet(server.NewServer)
)

func NewApp(
	db *gorm.DB,
	server *server.Server,
) *App {
	return &App{
		DB:     db,
		Server: server,
	}
}

type App struct {
	DB     *gorm.DB
	Server *server.Server
}

func (_this *App) Run() error {
	return _this.Server.Start()
}

func NewMySQLConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/terrorblade")
	checkError(err)
	return db
}

func NewRoutingFn(userHandler *handlers.UserHandler) server.RoutingFn {
	return func(router *echo.Echo) {
		v1 := router.Group("/api/v1")

		user := v1.Group("/users")
		{
			// user.POST("/", func(c echo.Context) error {
			// 	return c.JSON(200, map[string]interface{}{"Status": "OK"})
			// })
			user.POST("/", userHandler.CreateUser())
			user.GET("/:id", userHandler.GetUser())
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("An error occurred: %v", err)
	}
}
