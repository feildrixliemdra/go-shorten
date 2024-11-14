package util

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandomString generates a random alphabet-only string of a specified length
func RandomString(n int) string {
	// Seed random number generator to ensure different outputs
	rand.NewSource(time.Now().UnixNano())

	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		index := rand.Intn(len(letters))
		sb.WriteByte(letters[index])
	}
	return sb.String()
}

// ShortenURL generates a unique shortened alphabet-only string for a given URL
func ShortenURL(url string) string {
	// Append a random string to the URL to ensure uniqueness
	randomizedURL := fmt.Sprintf("%s%s", url, RandomString(6))

	hash := md5.Sum([]byte(randomizedURL))

	// Convert the hash to a base62 alphabet-only string
	shortURL := ""
	for _, b := range hash[:6] { // Use first 6 bytes to keep length short
		shortURL += string(letters[int(b)%len(letters)])
	}

	return shortURL
}

// RemoveNonAlphabets removes any character that is not an alphabet (A-Z, a-z).
func RemoveNonAlphabets(input string) string {
	result := ""
	for _, char := range input {
		if unicode.IsLetter(char) {
			result += string(char)
		}
	}
	return result
}
