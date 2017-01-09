package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte("123456"))
	cipherStr := md5Ctx.Sum([]byte("45"))
	fmt.Print(cipherStr)
	fmt.Print("\n")
	fmt.Print(hex.EncodeToString(cipherStr))
	fmt.Print("\n")
}
