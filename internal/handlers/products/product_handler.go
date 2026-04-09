package products

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/l10-bhushan/ecom_project/internal/services/products"
)

type Handler struct {
	service products.Service
}

func NewHandler(service products.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Fetch All the products
func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	// Call the service
	// Return Json encoded response
	err := h.service.GetAllProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	products := struct {
		Products []string `json:"products"`
	}{
		Products: []string{"Apple", "Mango"},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Create a product
func (h *Handler) CreateProduct() {}
