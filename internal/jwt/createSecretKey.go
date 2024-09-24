package jwtpackage

import (
	"iter"
	"math/rand"
	"strconv"
)

var secretKey []byte

func addRandomNumberToSecreteKey() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range 32 {
			i = rand.Int()
			if !yield(i) {
				return
			}
		}
	}
}

func CreateSecretKey() {
	var val int
	for i := range addRandomNumberToSecreteKey() {
		val += i
	}
	secretKey = []byte(strconv.Itoa(val))
}
