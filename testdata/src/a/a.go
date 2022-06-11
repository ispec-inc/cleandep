package a

import (
	"crypto/rand"
	"fmt" // want "package 'a' cannot depend on package 'fmt'"
	"math/big"
	mathrand "math/rand" // want "package 'a' cannot depend on package 'math/rand'"
	"net/http"           // want "package 'a' cannot depend on package 'net/http'"
	"net/rpc"            // want "package 'a' cannot depend on package 'net/rpc'"
)

func f() {
	fmt.Println(mathrand.Int())
	fmt.Println(rand.Int(rand.Reader, big.NewInt(100)))
	var _ = http.DefaultClient
	var _ = rpc.Client{}
}
