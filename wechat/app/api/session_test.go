package api

import (
	"log"
	"testing"
)

func TestSession(t *testing.T) {
	session, err := Session("", "", "")
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v", session)
}
