package idgen

import (
	"encoding/base64"
	"strings"
	"sync"
	"testing"
	"time"
)

// ✅ Test: ID uniqueness
func TestGenerate_UniqueIDs(t *testing.T) {
	idGen := New(6, "", "")
	id1 := idGen.Generate()
	id2 := idGen.Generate()

	if id1 == id2 {
		t.Errorf("Generated IDs should be unique, got %s and %s", id1, id2)
	}
}

// ✅ Test: Correct ID length
func TestGenerate_Length(t *testing.T) {
	idGen := New(8, "1234567890ABCDEF", "") // 8-char random part
	id := idGen.Generate()

	// Check if ID length matches the expected format
	expectedLength := len(time.Now().Format("060102150405")) + 8 + 4 // timestamp + random part + counter
	if len(id) != expectedLength {
		t.Errorf("Generated ID has incorrect length: got %d, expected %d", len(id), expectedLength)
	}
}

// ✅ Test: Correct timestamp format
func TestGenerate_TimestampFormat(t *testing.T) {
	timeFormat := "20060102150405" // YYYYMMDDHHMMSS
	idGen := New(6, "", timeFormat)
	id := idGen.Generate()

	timestamp := id[:len(time.Now().Format(timeFormat))]
	if _, err := time.Parse(timeFormat, timestamp); err != nil {
		t.Errorf("Generated ID does not start with a valid timestamp: got %s", timestamp)
	}
}

// ✅ Test: Thread safety (concurrent ID generation)
func TestGenerate_Concurrency(t *testing.T) {
	idGen := New(6, "", "")
	numWorkers := 100
	ids := make(map[string]bool)
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			id := idGen.Generate()

			mu.Lock()
			if ids[id] {
				t.Errorf("Duplicate ID detected: %s", id)
			}
			ids[id] = true
			mu.Unlock()
		}()
	}
	wg.Wait()
}

// ✅ Test: Base64 encoding
func TestEncodeBase64(t *testing.T) {
	id := "240210120530XYZ9876"
	encoded := EncodeBase64(id)

	// Check if it is a valid Base64 string
	_, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		t.Errorf("Base64 encoding failed: %v", err)
	}
}

// ✅ Test: Base58 encoding
func TestEncodeBase58(t *testing.T) {
	id := "240210120530XYZ9876"
	encoded := EncodeBase58(id)

	// Ensure it contains only Base58 characters
	const base58Chars = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	for _, ch := range encoded {
		if !strings.ContainsRune(base58Chars, ch) {
			t.Errorf("Base58 encoding contains invalid character: %c", ch)
		}
	}
}
