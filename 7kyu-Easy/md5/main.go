package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Print(PassHash("Serhat"))
}

func PassHash(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
