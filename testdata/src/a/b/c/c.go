package c

import (
	"crypto/rand"
	"log" // want "package 'c' cannot depend on package 'log'"
	"math/big"
	mathrand "math/rand" // want "package 'c' cannot depend on package 'math/rand'"
	"net/http"           // want "package 'c' cannot depend on package 'net/http'"
	"net/netip"          // want "package 'c' cannot depend on package 'net/netip'"
)

func f() {
	log.Println(mathrand.Int())
	log.Println(rand.Int(rand.Reader, big.NewInt(100)))
	var _ = http.DefaultClient
	var _ = netip.Addr{}
}
