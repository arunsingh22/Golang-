package main

import (
	"crypto/sha512"
	"fmt"
	"io"
)

func SecureHashing() {
	h := sha512.New()
	io.WriteString(h, "hello")
	sign := h.Sum(nil)
	fmt.Println(sign)
}
