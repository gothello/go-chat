package utils

import (

	"time"
	"math/rand"
)

func GetRandom() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}

func GetRandomInt64() int64 {
	return GetRandom().Int63()
}