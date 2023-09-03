package encoding

import (
	"encoding/json"
	"net/http"
)

func Encode(w http.ResponseWriter, body interface{}) error {
	err := json.NewEncoder(w).Encode(&body)
	if err != nil {
		return err
	}
	return nil
}

func Decode(r *http.Request, body interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return err
	}
	return nil
}
