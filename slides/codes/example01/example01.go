package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func randStr(length int) string { // OMIT START
	buf := make([]byte, length)
	rand.Read(buf) // HLexample
	return base64.StdEncoding.EncodeToString(buf)
} // OMIT END

func main() {
	l := 32
	fmt.Println(randStr(l))
}
