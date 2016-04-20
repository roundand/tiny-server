package main

import (
	"encoding/json"
	"fmt"
  "log"
	"net/http"
)

type UserId struct {
	UserId     string
	IsNewUser  bool
	IsVerified bool
}

var user = UserId{
	UserId:     "4d27dee1-9630-49cb-bed1-c656e160fddd",
	IsNewUser:  false,
	IsVerified: true,
}

var userJson []byte

func handler(w http.ResponseWriter, r *http.Request) {
  userJson, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}
	fmt.Fprintf(w, "here's some json >>%s<<", userJson)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
