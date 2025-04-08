package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/zeromicro/go-zero/core/logx"
	"regexp"
)

func InList(list []string, key string) (ok bool) {
	for _, s := range list {
		if s == key {
			return true
		}
	}
	return false
}

func InListByRegex(list []string, key string) (ok bool) {
	for _, s := range list {
		regex, err := regexp.Compile(s)
		if err != nil {
			logx.Error(err)
			return
		}
		if regex.MatchString(key) {
			return true
		}
	}
	return false
}
func MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
