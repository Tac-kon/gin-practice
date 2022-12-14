// product_dto.go

package products

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/gouser/api/utils/errors"
)

// Product - defines product info uploaded by user
type Product struct {
	gorm.Model
	// ID        uint64 `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Price  uint64 `json:"price"`
	Img    []byte `json:"img"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt time.Time
}

// Validate - check parameters user inputs
func (p *Product) Validate() *errors.ApiErr {
	p.Name = strings.TrimSpace(strings.ToLower(p.Name))
	if p.Name == "" {
		return errors.NewBadRequestError("invalid product name")
	}
	return nil
}