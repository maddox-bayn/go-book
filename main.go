/*
{
  "id": 1,
  "email": "user@example.com",
  "roles": ["admin", "editor"],
  "profile": {
    "age": 30,
    "bio": "Go dev"
  }
}
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	Age int    `json:"age"`
	Bio string `json:"bio"`
}

type User struct {
	ID      int      `json:"id"`
	Email   string   `json:"email"`
	Role    []string `json:"role"`
	Profile Profile  `json:"profile"`
}

func main() {
	raw := []byte(`{
	"id": 1,
	"email": "maddoxgmail.com",
	"role": ["admin", "master"],
	"profile": {"age": 30, "bio": "Go dev"}
	}`)
	var u User
	err := json.Unmarshal(raw, &u)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
}
