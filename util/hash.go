package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

func MD5HashString(data string) (string, error) {
	md5Hash:=md5.New()
	md5Hash.Write([]byte(data))
	passwordBytes:=md5Hash.Sum(nil)
	return hex.EncodeToString(passwordBytes), nil
}

func MD5HashReader(reader io.Reader) (string,error)  {
	bytes,err:=ioutil.ReadAll(reader)
	if err!=nil {
		fmt.Println(err.Error())
		return "",err
	}
	md5Hash:=md5.New()
	md5Hash.Write(bytes)
	hashBytes:=md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}

func SHA256Hash(data []byte) ([]byte) {
	sha256Hash:=sha256.New()
	sha256Hash.Write(data)
	return sha256Hash.Sum(nil)
}