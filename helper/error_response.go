package helper

import (
	"github.com/AsrofunNiam/technical-test-credit-plus/model/web"
	"github.com/gin-gonic/gin"
)

func SendErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, web.WebResponse{
		Success: false,
		Message: message,
	})
	c.Next()
}
