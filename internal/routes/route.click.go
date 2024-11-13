package routes

import (
	"click-counter/internal/handlers"
	"click-counter/internal/repositories/postgre"
	"click-counter/internal/services"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type ClickCounter struct {
	clickHandler *handlers.ClickHandler
	e            *echo.Group
}

func NewClickRouter(
	conn *sqlx.DB,
	e *echo.Group,
) *ClickCounter {
	clickRepository := postgre.NewPostgreClickRepository(conn)
	clickService := services.NewClickService(clickRepository)
	clickHandler := handlers.NewClickHandler(clickService)

	return &ClickCounter{
		clickHandler: clickHandler,
		e:            e,
	}
}

func (r *ClickCounter) Register() {
	tokenGroup := r.e.Group("")

	tokenGroup.GET("/counter/:banner_id", r.clickHandler.Generate)
	tokenGroup.POST("/stats/:banner_id", r.clickHandler.GetStats)
}
