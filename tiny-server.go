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

// newUUID generates a random UUID according to RFC 4122
// (nicked from http://play.golang.org/p/4FkNSiUDMg, via http://stackoverflow.com/questions/15130321/is-there-a-method-to-generate-a-uuid-with-go-language#comment42045723_15130965)
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

func ReportError(w http.ResponseWriter, err error, code int) {
	http.Error(w, err.Error(), code)
	log.Printf("Error %v: %s\n", code, err.Error())
}

func handler(w http.ResponseWriter, r *http.Request) {

	uuid, err := newUUID()
	if err != nil {
		ReportError(w, err, 500)
		return
	}

	user := UserId{
		UserId:     uuid,
		IsNewUser:  false,
		IsVerified: true,
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		ReportError(w, err, 500)
		return
	}
	fmt.Fprintf(w, "%s", userJson)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
