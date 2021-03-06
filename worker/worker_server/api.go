package rest_server

import ( 
  "net/http" 
  "github.com/hyper-ml/hyperml/server/pkg/base"
 
)

func raiseError(error_msg string) error{
  return base.HTTPErrorf(http.StatusInternalServerError, error_msg)
}

func (h *Handler) handleRoot() error {
  response := map[string]interface{}{
    "workhorse": "version 0.1",
  }

  h.writeJSON(response)
  return nil
}