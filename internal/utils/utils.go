package utils

import "math/rand"

func GenerateNumberRandom() int {
	return rand.Intn(99) + 1
}

func GetNumberRandomWithLimit(limit int) int {
	return rand.Intn(limit-1) + 1
}
