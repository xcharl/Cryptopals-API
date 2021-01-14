package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"./algorithms"
)

func main() {
	fmt.Println("hello 1")
	fmt.Println("hello again!!!!!")

	base64.StdEncoding.EncodeToString([]byte("hello"))

	x, err := hex.DecodeString("ff0d44en")
	y, _ := algorithms.EncodeToHex(x)

	fmt.Println(err)
	fmt.Println(y)
}
