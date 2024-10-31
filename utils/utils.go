package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)
 
func ParseJSON(r *http.Request, payload any) error {
  if r.Body == nil {
    return fmt.Errorf("missing Request body") 
  } 
    
  return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
  23:18
}
