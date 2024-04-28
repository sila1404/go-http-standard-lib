package product

import (
	"net/http"

	"github.com/sila1404/go-http-standard-lib/types"
	"github.com/sila1404/go-http-standard-lib/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoute(router *http.ServeMux) {
	router.HandleFunc("GET /products", h.handleCreateProduct)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}
