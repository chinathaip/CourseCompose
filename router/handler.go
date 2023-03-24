package router

import (
	"net/http"

	"github.com/chinathaip/coursecompose/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type Handler struct {
	mongoService *service.MongoService
}

func NewHandler(mongoService *service.MongoService) *Handler {
	return &Handler{mongoService: mongoService}
}

func (h *Handler) HandleGetAllCourses(c echo.Context) error {
	var filter bson.M
	switch c.Param("program") {
	case "it":
		filter = bson.M{"programs": bson.M{"$regex": "IT"}}
	default:
		filter = bson.M{"programs": "Yo"}
	}

	courses := h.mongoService.GetAllCourses(filter)

	if len(courses) <= 0 {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, courses)
}
