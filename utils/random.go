package utils

import "math/rand"

var runes = []rune("0123456789abcdefghijklmnopgrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomURL(size int) string {
	strArr := make([]rune, size)
	for i := range strArr {
		strArr[i] = runes[rand.Intn(len(runes))]
	}
	return string(strArr)
}
