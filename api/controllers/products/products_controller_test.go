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
	"github.com/jinzhu/gorm"
	"github.com/gouser/api/domain/products"
	"github.com/gouser/api/utils/errors"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductNoError(t *testing.T) {
	// Arrange ---
	p := products.Product{ // 修正
		Model: gorm.Model{ID: 123},
		Name:  "coca cola",
	}
	c, response := requestHandler(p)
	// Act ---
	CreateProduct(c)

	// Assert ---
	var product products.Product
	err := json.Unmarshal(response.Body.Bytes(), &product)
	assert.EqualValues(t, http.StatusCreated, response.Code)
	assert.Nil(t, err)
	fmt.Println(product)
	assert.EqualValues(t, uint64(123), product.ID)
}

func TestGetProductNoError(t *testing.T) {
	// Arrange
	p := products.Product{ // 修正
		Model: gorm.Model{ID: 1},
		Name:  "coca cola",
	}
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
