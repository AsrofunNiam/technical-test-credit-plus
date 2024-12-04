package exception

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/AsrofunNiam/technical-test-credit-plus/helper"
	"github.com/AsrofunNiam/technical-test-credit-plus/model/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(c *gin.Context, err interface{}) {
	if validationError(c, err) {
		return
	}

	if sendToResponseError(c, err) {
		return
	}

	if permissionDeniedError(c, err) {
		return
	}

	if foreignKeyError(c, err) {
		return
	}

	if recordNotFoundError(c, err) {
		return
	}

	if unauthorizedError(c, err) {
		return
	}

	if duplicateError(c, err) {
		return
	}

	internalServerError(c, err)
}

func validationError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		webResponse := web.WebResponse{
			Success: false,
			Message: "Bad Request",
			Data:    helper.ErrorRequestMessage(exception),
		}

		c.JSON(http.StatusBadRequest, webResponse)
		return true
	} else {
		return false
	}
}

func recordNotFoundError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok && exception.Error() == "record not found" {
		webResponse := web.WebResponse{
			Success: true,
			Message: "Record not found",
		}

		c.JSON(http.StatusOK, webResponse)
		return true
	}
	return false
}

func sendToResponseError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(*ErrorSendToResponse)
	if ok {
		webResponse := web.WebResponse{
			Success: false,
			Message: exception.Error(),
		}

		c.JSON(http.StatusBadRequest, webResponse)
		return true
	}
	return false
}

func unauthorizedError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok && (errors.Is(exception, ErrUnauthorized) || errors.Is(exception, ErrRefreshTokenExpired)) {
		webResponse := web.WebResponse{
			Success: false,
			Message: exception.Error(),
		}

		c.JSON(http.StatusUnauthorized, webResponse)
		return true
	}
	return false
}

func permissionDeniedError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok && errors.Is(exception, ErrPermissionDenied) {
		webResponse := web.WebResponse{
			Success: false,
			Message: exception.Error(),
		}
		c.JSON(http.StatusForbidden, webResponse)
		return true
	}
	return false
}

func duplicateError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok {
		if strings.Contains(exception.Error(), "Error 1062 (23000): Duplicate entry") {
			webResponse := web.WebResponse{
				Success: false,
				Message: helper.ErrorDuplicateMessage(exception),
			}

			c.JSON(http.StatusBadRequest, webResponse)
			return true
		}
	}
	return false
}

func foreignKeyError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok {
		if strings.Contains(exception.Error(), "Error 1451 (23000): Cannot delete or update a parent row") {
			webResponse := web.WebResponse{
				Success: false,
				Message: helper.ErrorForeignMessage(exception),
			}

			c.JSON(http.StatusBadRequest, webResponse)
			return true
		}
	}
	return false
}

func internalServerError(c *gin.Context, err interface{}) {
	log.Printf("Internal Server Error: %v", err)
	webResponse := web.WebResponse{
		Success: false,
		Message: "Internal Server Error",
	}
	c.JSON(http.StatusInternalServerError, webResponse)
}
