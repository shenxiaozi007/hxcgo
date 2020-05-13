package encrypt

import (
	"log"
	"testing"
)

func TestCrypt(t *testing.T) {
	rawPassword := "123456"
	pw, err := Password(rawPassword)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("len=%d,str=%s", len(pw), pw)
	log.Println(CompareHashAndPassword(pw, rawPassword))
}
