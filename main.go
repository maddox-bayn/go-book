package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	var u User

	json.Unmarshal([]byte(`{"id":1,"name":"maddox","email":"maddox@gmail.com"}`), &u)

	fmt.Println(u.ID, u.Name, u.Email)
}
