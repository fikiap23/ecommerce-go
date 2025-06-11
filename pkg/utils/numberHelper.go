package utils

import (
	"math/rand"
	"strconv"
	"time"
)

// GenRandomNumber generates a random integer with exactly 'length' digits.
func GenRandomNumber(length int) int {
	if length <= 0 {
		return 0
	}

	// Local random generator with time-based seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Ensure the first digit is not 0
	firstDigit := r.Intn(9) + 1 // 1-9
	numberStr := strconv.Itoa(firstDigit)

	for i := 1; i < length; i++ {
		digit := r.Intn(10) // 0-9
		numberStr += strconv.Itoa(digit)
	}

	result, _ := strconv.Atoi(numberStr)
	return result
}
