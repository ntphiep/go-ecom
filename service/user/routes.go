package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ntphiep/go-ecom/types"
	"github.com/ntphiep/go-ecom/utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}
          
func (h *Handler) RegisterRoutes(router *mux.Router) {
  router.HandleFunc("/login", h.handleLogin).Methods("POST")
  router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister (w http.ResponseWriter, r *http.Request) { 
	// get payload from request body
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
  	}

	// check if the user already exists
	
}
