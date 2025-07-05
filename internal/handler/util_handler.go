package handler

import (
	"encoding/json"
	"net/http"
)

func JsonDecode[T any](w http.ResponseWriter, r *http.Request, target *T) error {
	if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
		return err
	}
	return nil
}
