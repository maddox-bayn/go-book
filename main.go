package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id   int
	Name string
}

func handler(w http.ResponseWriter, r http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Bad request: error", 400)
		return
	}
	w.Header().Set("application", "json")
	json.NewEncoder(w).Encode(u)
}
