package handlers

import (
	"encoding/json"
	"github.com/DerryRenaldy/Todo-List-App/constants"
	"net/http"
)

// ResponseJson accept a writer and a response struct that want to be sent as a json response
func ResponseJson(w http.ResponseWriter, data interface{}) error {
	w.Header().Set(constants.HeaderContentType, constants.MIMEApplicationJson)
	return json.NewEncoder(w).Encode(data)
}
