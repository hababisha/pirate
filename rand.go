package main

import (
	"math/rand"
	"time"
)

func generateRand() int{
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	min := 1000
	max := 9999

	randomNumber := r.Intn(max-min) + min
	return randomNumber
}