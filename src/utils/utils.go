package utils

import (
	"math/rand"
	"time"
)

//RandomString - returns a random string;
func RandomString() string {

	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	//Min length = 80, max 120
	rand.Seed(time.Now().UnixNano())
	length := 80 + rand.Intn(120-80+1)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
