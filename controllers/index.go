package controllers

import (
	"encoding/json"
	"net/http"

	"../structs"
)

// IndexHandlerGET handler get for index request
func IndexHandlerGET(w http.ResponseWriter, r *http.Request) {
	var (
		responseStruct structs.ResponseStruct
	)
	result := &responseStruct
	w.Header().Set("Content-Type", "application/json")
	result.Status = true
	result.Result = map[string]interface{}{
		"connection": true,
	}
	json.NewEncoder(w).Encode(result)
}
