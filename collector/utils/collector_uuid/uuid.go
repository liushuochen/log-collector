// Package collector_uuid used to generate an uuid string.
package collector_uuid

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	upperLetterPool = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerLetterPool = "abcdefghijklmnopqrstuvwxyz"
	numberPool      = "0123456789"
)

// New function returns an uuid with lower letters and numbers.
// For example: 3483da85-47b8-429b-8346-9bfb9325c778
func New() string {
	parts := make([]string, 5)
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			parts[0] = generateRandomString(8, lowerLetterPool, numberPool)
		case 4:
			parts[4] = generateRandomString(12, lowerLetterPool, numberPool)
		default:
			parts[i] = generateRandomString(4, lowerLetterPool, numberPool)
		}
	}
	return strings.Join(parts, "-")
}


func generateRandomString(length int, pools ...string) string {
	resultList := make([]string, length, length)
	pool := strings.Join(pools, "")
	poolLength := int64(len(pool))

	for i := 0; i < length; i++ {
		random, _ := rand.Int(rand.Reader, big.NewInt(poolLength))
		resultList[i] = string(pool[int(random.Uint64())])
	}
	return strings.Join(resultList, "")
}
