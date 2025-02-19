package product

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Images      string `json:"img_url" validate:"required,url"`
}

func NewProduct(name string, description string, images string) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Images:      images,
	}
}
