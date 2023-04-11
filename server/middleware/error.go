package middleware

import (
	"encoding/json"
	"github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
	"net/http"
)

type ErrHandler func(http.ResponseWriter, *http.Request) error

func (fn ErrHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			xerr := cErrors.CustomError{
				Code:    500,
				Status:  "Error service panic",
				Message: "Something wrong is happening, service is panicking",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(xerr.Code)
			_ = json.NewEncoder(w).Encode(xerr)
			return
		}
	}()
	if err := fn(w, r); err != nil {
		xerr := err.(*cErrors.CustomError)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(xerr.Code)
		_ = json.NewEncoder(w).Encode(xerr)
		return
	}

}
