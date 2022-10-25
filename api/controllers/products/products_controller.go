// products_controller.go

package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gouser/api/domain/products"
	"github.com/gouser/api/services"
	"github.com/gouser/api/utils/errors"
)

// productIDの取得を外に出した
func getProductID(productIDParam string) (uint, *errors.ApiErr) {
	productID, productErr := strconv.ParseUint(productIDParam, 10, 64)
	if productErr != nil {
		return 0, errors.NewBadRequestError("product id should be a number")
	}
	return uint(productID), nil
}

// GetProduct - Get product by product id
func GetProduct(c *gin.Context) {
	productID, idErr := getProductID(c.Param("product_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	product, getErr := services.GetProduct(uint(productID))
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct - Create product
func CreateProduct(c *gin.Context) {
	var product products.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}
	newProduct, saveErr := services.CreateProduct(product)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}

// UpdateProduct - Update product
func UpdateProduct(c *gin.Context) {
	productID, idErr := getProductID(c.Param("product_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var product products.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}

	product.ID = productID

	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UpdateProduct(isPartial, product)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteProduct - Delete product
func DeleteProduct(c *gin.Context) {
	productID, idErr := getProductID(c.Param("product_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteProduct(productID); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}