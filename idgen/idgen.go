package idgen

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	defaultCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	defaultLength  = 6
	defaultTimeFmt = "060102150405"
	minCharsetLen  = 2
)

var (
	globalRand   *rand.Rand
	globalRandMu sync.Mutex
	randOnce     sync.Once
)

func initRand() {
	src := rand.NewSource(time.Now().UnixNano())
	globalRand = rand.New(src)
}

type ShortIDGenerator struct {
	length     int
	charset    string
	timeFormat string
	counter    uint64
}

func New(length int, charset, timeFormat string) *ShortIDGenerator {
	if length <= 0 {
		length = defaultLength
	}
	if charset == "" || len(charset) < minCharsetLen {
		charset = defaultCharset
	}
	if timeFormat == "" {
		timeFormat = defaultTimeFmt
	}

	randOnce.Do(initRand)

	return &ShortIDGenerator{
		length:     length,
		charset:    charset,
		timeFormat: timeFormat,
	}
}

func (g *ShortIDGenerator) Generate() string {
	timestamp := time.Now().Format(g.timeFormat)
	randomStr := generateRandomString(g.length, g.charset)
	counter := atomic.AddUint64(&g.counter, 1) % 10000
	return fmt.Sprintf("%s%s%04d", timestamp, randomStr, counter)
}

func generateRandomString(length int, charset string) string {
	globalRandMu.Lock()
	defer globalRandMu.Unlock()

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[globalRand.Intn(len(charset))]
	}
	return string(result)
}

func EncodeBase64(input string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(input))
}

func EncodeBase58(input string) string {
	if len(input) == 0 {
		return ""
	}

	const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	num := new(big.Int).SetBytes([]byte(input))
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := new(big.Int)

	var result []byte
	for num.Cmp(zero) > 0 {
		num.DivMod(num, base, mod)
		result = append([]byte{alphabet[mod.Int64()]}, result...)
	}

	return string(result)
}
