package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type testHandler struct {
	testHandlerService *service
}

func NewHandler(v1 *gin.RouterGroup, service *service) {
	handler := &testHandler{service}

	main := v1.Group("test")

	main.GET("/:string", handler.GetString)
}

func (h *testHandler) GetString(c *gin.Context) {
	textValue := c.Param("string")

	test := h.testHandlerService.GetTest(textValue)

	c.String(http.StatusOK, test)
}
