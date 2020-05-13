package cache

import (
	"bytes"
	"encoding/gob"
	"log"
	"testing"
)

type Person struct {
	ID   uint
	Name string
	Age  int
}

func TestCache(t *testing.T) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	alan := &Person{1, "alan", 23}
	err := enc.Encode(alan)

	if err != nil {
		log.Fatalln("encode error: ", err)
	}

	dec := gob.NewDecoder(&buf)

	people := &Person{}
	err = dec.Decode(people)
	if err != nil {
		log.Fatalln("encode error: ", err)
	}

	log.Printf("%v", people)
}
