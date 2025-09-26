package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"vetsys/internal/domain"

	"vetsys/internal/usecase/customer"

	"github.com/go-chi/chi/v5"
)

// CustomerHandler is the HTTP adapter.
// It receives the Usecase as a dependency.
type CustomerHandler struct {
	CreateUC  *customer.CreateCustomerUsecase
	GetByIDUC *customer.GetCustomerByIDUsecase
}

// NewCustomerHandler is the constructor.
func NewCustomerHandler(createUC *customer.CreateCustomerUsecase, getByIDUC *customer.GetCustomerByIDUsecase) *CustomerHandler {
	return &CustomerHandler{
		CreateUC:  createUC,
		GetByIDUC: getByIDUC,
	}
}

// RegisterRoutes configures the routes for the Handler (example with Chi).
func (h *CustomerHandler) RegisterRoutes(r chi.Router) {
	r.Post("/customers", h.CreateCustomer)
	r.Get("/customers/{id}", h.GetCustomerByID)
}

// ----------------------------------------------------------------------
// POST /customers
// ----------------------------------------------------------------------

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var req CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 1. Translate the Request DTO (HTTP) to the Usecase Input DTO
	input := customer.CreateCustomerInput{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	// 2. Execute the Usecase (Business Logic)
	output, err := h.CreateUC.Execute(r.Context(), input)
	if err != nil {
		// Handle domain-specific errors here if necessary (e.g. 409 Conflict)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Return the response (using the Usecase Output DTO)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(output)
}

// ----------------------------------------------------------------------
// GET /customers/{id}
// ----------------------------------------------------------------------

func (h *CustomerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// 1. Executa o Usecase
	output, err := h.GetByIDUC.Execute(r.Context(), id)
	if err != nil {
		// Example: if it is a "not found" error (coming from Repository/Domain)
		if errors.Is(err, domain.ErrCustomerNotFound) {
			http.Error(w, "Customer not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Return the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
