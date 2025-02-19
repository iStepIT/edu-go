package product

type ProductCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Images      string `json:"img_url" validate:"required,url"`
}

type ProductUpdateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Images      string `json:"img_url" validate:"required,url"`
}
