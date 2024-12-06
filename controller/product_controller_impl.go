package controller

import (
	"net/http"

	"github.com/AsrofunNiam/technical-test-credit-plus/auth"
	"github.com/AsrofunNiam/technical-test-credit-plus/helper"
	"github.com/AsrofunNiam/technical-test-credit-plus/model/web"
	"github.com/AsrofunNiam/technical-test-credit-plus/service"
	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "name.like", "id.eq")
	productResponses := controller.ProductService.FindAll(auth, &filters, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(productResponses),
		Data:    productResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ProductControllerImpl) FindByID(c *gin.Context, auth *auth.AccessDetails) {
	productIDParam := c.Param("id")
	productID := helper.StringToUint(productIDParam)

	productResponse := controller.ProductService.FindByID(auth, &productID, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(productResponse),
		Data:    productResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}
