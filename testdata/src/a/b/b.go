package b

import (
	"crypto/rand"
	"log" // want "package 'a/b' cannot depend on package 'log'"
	"math/big"
	mathrand "math/rand" // want "package 'a/b' cannot depend on package 'math/rand'"
	"net/http"           // want "package 'a/b' cannot depend on package 'net/http'"
	"net/rpc"            // want "package 'a/b' cannot depend on package 'net/rpc'"
)

func f() {
	log.Println(mathrand.Int())
	log.Println(rand.Int(rand.Reader, big.NewInt(100)))
	var _ = http.DefaultClient
	var _ = rpc.Client{}
}
