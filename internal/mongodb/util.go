package mongodb

import (
	"encoding/base64"
	"fmt"
	"os"
)

// initFileFromBase64String ...
func initFileFromBase64String(filename, value string) (*os.File, error) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("mongodb.initFileFromBase64String - err: ", err)
		return nil, err
	}
	b := base64DecodeToBytes(value)
	if _, err := f.Write(b); err != nil {
		fmt.Println("mongodb.initFileFromBase64String - write file err: ", err)
		return nil, err
	}
	f.Sync()
	return f, nil
}

func base64DecodeToBytes(text string) []byte {
	s, _ := base64.StdEncoding.DecodeString(text)
	return s
}

func base64DecodeToString(text string) string {
	return string(base64DecodeToBytes(text))
}
