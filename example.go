package main

import (
	"fmt"
	"github.com/fxamacker/cbor/v2"
	"os"
)

func main() {
	f, err := os.OpenFile("output-go.cbor", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	v := [2]string{"hello", "world"}
	encoder := cbor.NewEncoder(f)
	err = encoder.Encode(v)
	if err != nil {
		panic(err)
	}

	fmt.Println(v)
}
