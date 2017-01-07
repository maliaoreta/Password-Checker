package routes

import (
	"encoding/json"
	"net/http"
)

func TestApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	test := make(map[string]bool)
	test["success"] = true

	j, _ := json.Marshal(test)
	w.Write(j)
}