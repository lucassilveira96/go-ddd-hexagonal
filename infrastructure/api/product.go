package api

import (
	"encoding/json"
	"go-ddd-hexagonal/app/service"
	"go-ddd-hexagonal/domain/model"
	"net/http"
	"strconv"
	"strings"
)

// RegisterProductRoutes configura as rotas para o gerenciamento de produtos
func RegisterProductRoutes(router *http.ServeMux, productService *service.ProductService) {
	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getAllProducts(w, r, productService)
		case http.MethodPost:
			createProduct(w, r, productService)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	router.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIDFromURL(r.URL.Path)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			getProductByID(w, r, productService, id)
		case http.MethodPut:
			updateProduct(w, r, productService, id)
		case http.MethodDelete:
			deleteProduct(w, r, productService, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func getAllProducts(w http.ResponseWriter, r *http.Request, productService *service.ProductService) {
	products, err := productService.FindAllProduct(r.Context())
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func createProduct(w http.ResponseWriter, r *http.Request, productService *service.ProductService) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if err := productService.CreateProduct(r.Context(), &product); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func getProductByID(w http.ResponseWriter, r *http.Request, productService *service.ProductService, id int64) {
	product, err := productService.FindByIdProduct(r.Context(), id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func updateProduct(w http.ResponseWriter, r *http.Request, productService *service.ProductService, id int64) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	product.ID = id
	if err := productService.UpdateProduct(r.Context(), &product); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request, productService *service.ProductService, id int64) {
	if err := productService.DeleteProduct(r.Context(), id); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// parseIDFromURL extrai o ID do produto a partir da URL
func parseIDFromURL(path string) (int64, error) {
	idStr := strings.TrimPrefix(path, "/products/")
	return strconv.ParseInt(idStr, 10, 64)
}
