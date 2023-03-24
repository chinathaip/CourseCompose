package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h *Handler) RegisterRoute() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/:program/courses", h.HandleGetAllCourses)

	return e
}
