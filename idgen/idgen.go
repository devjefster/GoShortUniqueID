package idgen

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Default values
const (
	defaultCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	defaultLength  = 6
	defaultTimeFmt = "060102150405" // YYMMDDHHmmss (default)
)

// ShortIDGenerator struct for configuration
type ShortIDGenerator struct {
	length     int
	charset    string
	timeFormat string
	counter    uint64
}

// New creates a new ShortIDGenerator with custom options
func New(length int, charset, timeFormat string) *ShortIDGenerator {
	if length <= 0 {
		length = defaultLength
	}
	if charset == "" {
		charset = defaultCharset
	}
	if timeFormat == "" {
		timeFormat = defaultTimeFmt
	}

	return &ShortIDGenerator{
		length:     length,
		charset:    charset,
		timeFormat: timeFormat,
	}
}

// Generate creates a unique, human-readable, timestamped ID
func (g *ShortIDGenerator) Generate() string {
	// Generate timestamp in the specified format
	timestamp := time.Now().Format(g.timeFormat)

	// Generate a random string
	randomStr := generateRandomString(g.length, g.charset)

	// Ensure uniqueness by adding an atomic counter
	counter := atomic.AddUint64(&g.counter, 1) % 10000 // Max 4-digit counter

	// Return combined ID
	return fmt.Sprintf("%s%s%04d", timestamp, randomStr, counter)
}

// generateRandomString creates a random string of given length
func generateRandomString(length int, charset string) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}
	return string(result)
}

// EncodeBase64 encodes a unique ID using Base64 (for URL-safe versions)
func EncodeBase64(input string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(input))
}

// EncodeBase58 encodes a unique ID using Base58 (shorter than Base64)
func EncodeBase58(input string) string {
	const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	var result string
	num := binary.BigEndian.Uint64([]byte(input))

	for num > 0 {
		remainder := num % 58
		num = num / 58
		result = string(alphabet[remainder]) + result
	}
	return result
}
