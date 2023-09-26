package JSON

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Parse(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(&data)
}


func ErrorResponse(w http.ResponseWriter, status int, err string, message string) {
	errMessage := fmt.Sprintf("%s. %s", message, err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(errMessage))
}

func Response(w http.ResponseWriter, status int, data interface{}) {
	jsonResp, err := json.Marshal(data)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Error happened in JSON marshal.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(jsonResp))
}
