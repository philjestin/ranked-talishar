package test_util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const numbers = "1234567890"

func init() {
		rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
		return min + rand.Int63n(max - min + 1) // Random int
}

func RandomString(n int) string {
		var sb strings.Builder
		k := len(alphabet)

		for i := 0; i < n; i++ {
				c := alphabet[rand.Intn(k)]
				sb.WriteByte(c)
		}

		return sb.String()
}

func RandomNumberString(n int) string {
			var sb strings.Builder
		k := len(numbers)

		for i := 0; i < n; i++ {
				c := numbers[rand.Intn(k)]
				sb.WriteByte(c)
		}

		return sb.String()
}

func RandomFirstName() string {
		return RandomString(12)
}

func RandomLastName() string {
	return RandomString(12)
}

func RandomPhoneNumber() string {
		return RandomNumberString(10)
}

func RandomStreet() string {
		return RandomString(15)
}

