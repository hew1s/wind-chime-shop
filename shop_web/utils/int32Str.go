package utils

import "strconv"

func IntStr(str string) int32 {
	int, _ := strconv.ParseInt(str, 10, 64)
	re := int32(int)
	return re
}
