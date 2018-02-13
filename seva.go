package main

import (
	"math/rand"
	"time"
)

func gopnicAnswer() string {
	sr := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(sr)
	return gopnik[r1.Intn(len(gopnik))]
}
