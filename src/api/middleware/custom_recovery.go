package middleware

import (
	"net/http"

	"github.com/SajjadRezaei/go-clean-architecture/api/helper"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err any) {
	if err, ok := err.(error); ok {
		httpResponse := helper.GenerateBaseResponseWithError(nil, false, helper.CustomRecovery, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
		return
	}
	httpResponse := helper.GenerateBaseResponseWithAnyError(nil, false, helper.CustomRecovery, err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
}
