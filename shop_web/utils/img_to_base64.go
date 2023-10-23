package utils

import (
	"encoding/base64"
	"os"
)

func ImgToBase64(path string) string {
	file, _ := os.Open(path)
	defer file.Close()
	buf := make([]byte, 1000000)
	n, _ := file.Read(buf)
	return base64.StdEncoding.EncodeToString(buf[:n])
}
