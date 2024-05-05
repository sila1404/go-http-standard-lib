package cart

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sila1404/go-http-standard-lib/service/auth"
	"github.com/sila1404/go-http-standard-lib/types"
	"github.com/sila1404/go-http-standard-lib/utils"
)

type Handler struct {
	productStore types.ProductStore
	orderStore   types.OrderStore
	userStore    types.UserStore
}

func NewHandler(productStore types.ProductStore, orderStore types.OrderStore, userStore types.UserStore) *Handler {
	return &Handler{productStore: productStore, orderStore: orderStore, userStore: userStore}
}

func (h *Handler) RegisterRoute(router *http.ServeMux) {
	router.HandleFunc("POST /cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userStore))
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())

	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Get products
	productsID, err := getCartItemsID(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	product, err := h.productStore.GetProductsByID(productsID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(product, cart.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"total_price": totalPrice,
		"order_id":    orderID,
	})
}
