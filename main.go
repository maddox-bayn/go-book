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

	contaent, _ := json.Marshal(User{
		ID:    1,
		Name:  "Maddox",
		Email: "maddox@gmail.com",
	})
	fmt.Println(string(contaent))
}
