package controller

import (
	"encoding/json"
	"net/http"
)

type Res struct {
	Code    uint16
	Message string
	Data    interface{}
}

func ApiResponse(w http.ResponseWriter, res *Res) {
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(&res)
}
