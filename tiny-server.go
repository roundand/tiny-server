// tiny-server generates dummy JSON responses
package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type UserId struct {
	UserId     string
	IsNewUser  bool
	IsVerified bool
}

var userJson []byte

// newUUID generates a random UUID according to RFC 4122
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func handler(w http.ResponseWriter, r *http.Request) {

	uuid, err := newUUID()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	var user = UserId{
		UserId:     uuid,
		IsNewUser:  false,
		IsVerified: true,
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}
	fmt.Fprintf(w, "%s", userJson)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
