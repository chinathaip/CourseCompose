package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chinathaip/coursecompose/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestHandleGetAllCourse(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		expectFilter bson.M
		code         int
	}{
		{
			name:         "200 - Should have query filter as IT when path is IT",
			path:         "it",
			expectFilter: bson.M{"programs": bson.M{"$regex": "IT"}},
			code:         http.StatusOK,
		},
		{
			name:         "404 - Should have invalid query filter when path is not IT",
			path:         "some invalid path",
			expectFilter: bson.M{"programs": "Yo"},
			code:         http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var (
				req            = httptest.NewRequest(http.MethodGet, "/", nil)
				rec            = httptest.NewRecorder()
				c              = echo.New().NewContext(req, rec)
				mservice       = &mockMongoService{}
				h              = NewHandler(mservice)
				expectedFilter = test.expectFilter
			)
			c.SetParamNames("program")
			c.SetParamValues(test.path)
			c.SetPath("/courses")

			err := h.HandleGetAllCourses(c)

			assert.Equal(t, expectedFilter, mservice.filter)
			if err != nil {
				assert.Equal(t, echo.ErrNotFound, err)
			} else {
				assert.Equal(t, test.code, rec.Code)
			}
		})
	}
}

type mockMongoService struct {
	filter bson.M
}

func (s *mockMongoService) GetAllCourses(filter bson.M) []model.AllCourseResponse {
	s.filter = filter

	if filter["programs"] == "Yo" {
		return []model.AllCourseResponse{}
	}

	return []model.AllCourseResponse{
		{CourseCode: "ITE123"},
		{CourseCode: "ITE456"},
	}
}
