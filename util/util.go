package util

import (
	"crypto/sha1"
	"fmt"
)

func Sha1(bytes []byte)string{
	checkSum:=sha1.Sum(bytes)
	return fmt.Sprintf("%x",checkSum)
}
