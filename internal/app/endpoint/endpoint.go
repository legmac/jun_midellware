package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint { // констркутор
	return &Endpoint{
		s: s,
	} // созадем эндпоинт и возвращаем
}

func (e *Endpoint) Status(ctx echo.Context) error {
	d := e.s.DaysLeft()

	s := fmt.Sprintf("Days laeft: %d", d)

	err := ctx.String(http.StatusOK, s) //"test")
	if err != nil {
		return err
	}
	return nil
}
