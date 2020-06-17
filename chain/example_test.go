package chain

import (
	"io/ioutil"
	"log"
	"strings"
)

func ExampleJoinReader() {
	r1 := strings.NewReader("hello")
	r2 := strings.NewReader(",")
	r3 := strings.NewReader("world")
	r4 := strings.NewReader("!")

	r := JoinReader(r1, r2, r3, r4)
	defer r.Close()

	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Panicf("Read error: %s", err.Error())
	}
	log.Printf("Content: %s", string(data))
}
