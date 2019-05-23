package main

import (
	"fmt"
	"golang.org/x/crypto/blake2b"
	"strings"
)

func blakeHash(s ...string) string {
	joined := strings.Join(s, "|")
	bytes := blake2b.Sum512([]byte(joined))
	outHex := fmt.Sprintf("%x", bytes)
	return outHex
}
