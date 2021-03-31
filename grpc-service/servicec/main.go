package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))

	e.GET("/", func(c echo.Context) error {
		ctx := c.Request().Context()
		e.Logger.Infof("request context: %+v", ctx)

		fmt.Println("----Service C consume 30ms----")
		time.Sleep(time.Duration(30) * time.Millisecond)
		select {
		case <-ctx.Done():
			fmt.Println("[ServiceC] client has cancel: ", ctx.Err())
		}
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Info("ServiceC started at :8082")
	e.Logger.Fatal(e.Start(":8082"))
}
