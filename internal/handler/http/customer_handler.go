package http

import (
	"encoding/json"
	"net/http"

	"vetsys/internal/usecase/customer"

	"github.com/go-chi/chi/v5"
)

// CustomerHandler é o adaptador HTTP.
// Ele recebe o Usecase como dependência.
type CustomerHandler struct {
	CreateUC  *customer.CreateCustomerUsecase
	GetByIDUC *customer.GetCustomerByIDUsecase
}

// NewCustomerHandler é o construtor.
func NewCustomerHandler(createUC *customer.CreateCustomerUsecase, getByIDUC *customer.GetCustomerByIDUsecase) *CustomerHandler {
	return &CustomerHandler{
		CreateUC:  createUC,
		GetByIDUC: getByIDUC,
	}
}

// RegisterRoutes configura as rotas para o Handler (exemplo com Chi).
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

	// 1. Traduz o DTO de Requisição (HTTP) para o DTO de Entrada do Usecase
	input := customer.CreateCustomerInput{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	// 2. Executa o Usecase (Lógica de Negócio)
	output, err := h.CreateUC.Execute(r.Context(), input)
	if err != nil {
		// Tratar erros específicos do domínio aqui, se necessário (ex: 409 Conflict)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Retorna a resposta (usando o DTO de Saída do Usecase)
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
		// Exemplo: se for um erro de "não encontrado" (vindo do Repositório/Domínio)
		// if errors.Is(err, domain.ErrCustomerNotFound) {
		// 	http.Error(w, "Customer not found", http.StatusNotFound)
		// 	return
		// }
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Retorna a resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
