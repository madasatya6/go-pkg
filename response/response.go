package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, p interface{}, status int){
	
	ubahkeByte, err := json.Marshal(p)
	
	if err != nil {
		http.Error(w, "Error om", http.StatusBadRequest)
	}
	
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	w.Write([]byte(ubahkeByte))
}