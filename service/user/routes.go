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
  var payload types.RegisterUserPayload
  if err := utils.ParseJSON(r.Body, payload); err != nil {
    
  }
}
