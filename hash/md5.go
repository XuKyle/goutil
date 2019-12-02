package hash

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

func FileMd5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", errors.New(fmt.Sprintf("md5.go hash.FileMd5 os open file err %v", err))
	}
	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", errors.New(fmt.Sprintf("md5.go hash.FileMd5 io copy err %v", err))
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func StringMd5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}
