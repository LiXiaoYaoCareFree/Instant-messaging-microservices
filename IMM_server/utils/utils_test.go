package utils

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	md := MD5([]byte("1234"))
	fmt.Println(md)
}
func TestGetFilePrefix(t *testing.T) {
	fmt.Println(GetFilePrefix("name.png"))
	fmt.Println(GetFilePrefix("1.name.png"))
	fmt.Println(GetFilePrefix("1.aaa.name.png"))
}
