package product

import (
	"4-order-api/pkg/req"
	"4-order-api/pkg/res"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ProductHandlerDeps struct {
	ProductRepository *ProductRepository
}
type ProductHandler struct {
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		ProductRepository: deps.ProductRepository,
	}
	router.HandleFunc("POST /product", handler.Create())
	router.HandleFunc("PATCH /product/{id}", handler.Update())
	router.HandleFunc("DELETE /product/{id}", handler.Delete())
	router.HandleFunc("GET /prod/{id}", handler.GetById())
}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductCreateRequest](&w, r)
		if err != nil {
			return
		}
		newProduct := NewProduct(body.Name, body.Description, body.Images)

		createdProduct, err := handler.ProductRepository.Create(newProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdProduct, 201)
	}
}

func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductUpdateRequest](&w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, _ := parseId(idString)
		product, err := handler.ProductRepository.Update(&Product{
			Model:       gorm.Model{ID: id},
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, product, 201)
	}
}

func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, _ := parseId(idString)
		body, err := handler.ProductRepository.GetById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = handler.ProductRepository.Delete(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := ProductUpdateResponse{
			Name:   body.Name,
			Status: "deleted",
		}
		res.Json(w, data, 200)
	}
}

func (handler *ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, _ := parseId(idString)
		product, err := handler.ProductRepository.GetById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		res.Json(w, product, 200)
	}
}

func parseId(idString string) (uint, error) {
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
