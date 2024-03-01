package main

import (
	"log"

	"gopkg.in/zabawaba99/firego.v1"
)

func main() {
	f := firego.New("firebase.url", nil)

	v := map[string]string{"foo": "bar"}
	if err := f.Set(v); err != nil {
		log.Fatalf("error when set: %v", err)
	}
}
