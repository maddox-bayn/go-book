package Book

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var Users = map[int]User{}
var NextID = 1

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Book.writeJSON encode error: %v", err)
	}
}

// UsersHandler handles POST /users and GET /users
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		u.ID = NextID
		NextID++
		Users[u.ID] = u

		writeJSON(w, http.StatusCreated, u)

	case http.MethodGet:
		list := make([]User, 0, len(Users))
		for _, u := range Users {
			list = append(list, u)
		}
		writeJSON(w, http.StatusOK, list)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// UserHandler handles requests for a specific user: PUT and DELETE
func UserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	_, ok := Users[id]
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodPut:
		var updated User
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		updated.ID = id
		Users[id] = updated
		writeJSON(w, http.StatusOK, updated)

	case http.MethodDelete:
		delete(Users, id)
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
