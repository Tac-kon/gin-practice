// products_controller_test.go

package products

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gouser/api/domain/products"
	"github.com/gouser/api/utils/errors"
	"github.com/stretchr/testify/assert"
)

func getRequestHandler(id string) (*gin.Context, *httptest.ResponseRecorder) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	param := gin.Param{Key: "product_id", Value: id}
	c.Params = gin.Params{param}

	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/products/:product_id",
		nil,
	)

	return c, response
}

// 正常系
func TestGetProductNoError(t *testing.T) {
	// Arrange
	p := products.Product{ID: 1, Name: "coca cola"}
	c, _ := requestHandler(p)
	CreateProduct(c)

	c2, response := getRequestHandler("1")

	// Act ---
	GetProduct(c2)

	// Assert ---
	var product products.Product
	err := json.Unmarshal(response.Body.Bytes(), &product)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.Nil(t, err)
	assert.EqualValues(t, uint64(1), product.ID)
}

// 不正なIDのテスト
func TestGetProductWithInvalidID(t *testing.T) {
	// Arrange
	c, response := getRequestHandler("a")

	// Act ---
	GetProduct(c)

	// Assert ---
	var apiErr errors.ApiErr
	json.Unmarshal(response.Body.Bytes(), &apiErr)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, apiErr.Message, "product id should be a number")
	assert.EqualValues(t, apiErr.Status, 400)
	assert.EqualValues(t, apiErr.Error, "bad_request")
}

// 指定したIDのプロダクトが存在しないテスト
func TestGetProductWithNoProduct(t *testing.T) {
	// Arrange ---
	c, response := getRequestHandler("10000")

	// Act ---
	GetProduct(c)

	// Assert ---
	var apiErr errors.ApiErr
	json.Unmarshal(response.Body.Bytes(), &apiErr)
	assert.EqualValues(t, http.StatusNotFound, response.Code)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, apiErr.Message, "product 10000 not found")
	assert.EqualValues(t, apiErr.Status, 404)
	assert.EqualValues(t, apiErr.Error, "not_found")
}
