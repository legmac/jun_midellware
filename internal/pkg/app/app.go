package app

import (
	"fmt"
	"jun_midellware/internal/app/endpoint"
	"jun_midellware/internal/app/mw"
	"jun_midellware/internal/app/service"
	"log"

	"github.com/labstack/echo"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}
	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.echo = echo.New()

	a.echo.Use(mw.RoleCheck)
	a.echo.GET("/status", a.e.Status) // ,MW

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server started")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
