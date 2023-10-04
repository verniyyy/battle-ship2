package battleship2

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/verniyyy/battle-ship2/lib"
)

type HTTPServer struct {
	e    *echo.Echo
	ip   string
	port int
}

func NewHTTPServer(port int) HTTPServer {
	e := echo.New()

	ip, err := lib.IPAddr()
	if err != nil {
		e.Logger.Fatal(err)
	}

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// set routings
	e.File("/", "public/index.html")
	e.GET("/system", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusCreated, struct {
			ProjectName string
			Version     string
		}{
			ProjectName: "battle-ship2",
			Version:     "undefined",
		})
	})
	api := e.Group("/api")
	api.POST("/user", CreateUser)
	api.GET("/user/:id", DescribeUser)
	api.POST("/inqueue/:id", InQueue)
	api.POST("/ismatched/:id", IsMatched)
	api.POST("/connectroom/:id", ConnectRoom)

	return HTTPServer{
		e:    e,
		ip:   ip,
		port: port,
	}
}

func (s HTTPServer) Serve() error {
	go MatchMaking(s.e.Logger)
	return s.e.Start(addr(s.ip, s.port))
}

func (s HTTPServer) Logger() echo.Logger {
	return s.e.Logger
}

func addr(ip string, port int) string {
	return fmt.Sprintf("%s:%d", ip, port)
}
