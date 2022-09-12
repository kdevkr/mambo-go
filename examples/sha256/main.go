package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	plain := "HelloWorld"
	hash := sha256.New()
	hash.Write([]byte(plain))
	hashed := hash.Sum(nil)
	fmt.Println("plain:", plain)
	fmt.Println("hashed:", hex.EncodeToString(hashed))
}
