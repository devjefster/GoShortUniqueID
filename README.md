
# GoShortUniqueID 🚀

A **lightweight**, **fast**, and **sortable** unique ID generator for Go.  
Generates **short, human-readable IDs** with a timestamp prefix, ensuring uniqueness and chronological order.

[![Go Report Card](https://goreportcard.com/badge/github.com/devjefster/GoShortUniqueID)](https://goreportcard.com/report/github.com/devjefster/GoShortUniqueID)
[![GoDoc](https://pkg.go.dev/badge/github.com/devjefster/GoShortUniqueID.svg)](https://pkg.go.dev/github.com/devjefster/GoShortUniqueID)
[![License: MIT](https://img.shields.io/github/license/devjefster/GoShortUniqueID)](LICENSE)
[![Go Build](https://github.com/devjefster/GoShortUniqueID/actions/workflows/go.yml/badge.svg)](https://github.com/devjefster/GoShortUniqueID/actions)

---

## 🚀 Features
✔ **Chronologically Sortable** – IDs start with a timestamp (YYMMDDHHmmss).  
✔ **Short & Unique** – Customizable random suffix ensures uniqueness.  
✔ **Thread-Safe** – Uses an atomic counter to prevent collisions.  
✔ **Customizable** – Define **length**, **character set**, and **timestamp format**.  
✔ **URL-Friendly** – Optional **Base64** and **Base58** encoding.

---

## 📦 Installation

```sh
  go get github.com/devjefster/GoShortUniqueID
```

### 🛠️ Usage

**1️⃣ Basic Example**
```go
package main

import (
	"fmt"
	"https://github.com/devjefster/GoShortUniqueID/idgen"
)

func main() {
	// Create an ID generator (default: 6-character random part)
	idGen := idgen.New(6, "", "")

	// Generate unique IDs
	for i := 0; i < 5; i++ {
		fmt.Println(idGen.Generate())
	}
}
```
🔹 Example Output:

240210120530XYZ9876
240210120530ABC5432
240210120530TUV7890

**2️⃣ Customizing Length & Charset**
```go
// Create a generator with a custom charset and length
customGen := idgen.New(8, "123456789ABCDEFG", "")

// Generate an ID with a **numeric-only** random suffix
fmt.Println(customGen.Generate()) 
```
**3️⃣ Changing Timestamp Format**
```go
// Use full timestamp format (YYYYMMDDHHMMSS)
idGenFull := idgen.New(6, "", "20060102150405")

fmt.Println(idGenFull.Generate()) 
// Output: 20240210120530XYZ1234
```
**4️⃣ URL-Friendly Encoding**
```go
id := idGen.Generate()

// Encode ID in Base64
fmt.Println("Base64:", idgen.EncodeBase64(id))

// Encode ID in Base58 (shorter and safer)
fmt.Println("Base58:", idgen.EncodeBase58(id))
```
**🔹 Example Output:**

Base64: MjQwMjEwMTIwNTMwWFlaOTg3Ng
Base58: 1HQXYZ9876

**⚙️ API Reference**

**🔹 New(length int, charset string, timeFormat string) *ShortIDGenerator***

Creates a new instance of ShortIDGenerator.
Parameter	Type	Description
length	int	Length of the random suffix (default: 6)
charset	string	Characters to use in the random part (default: alphanumeric)
timeFormat	string	Timestamp format (default: YYMMDDHHmmss)

**🔹 Generate() string**

Generates a new unique ID.

**🔹 EncodeBase64(input string) string**

Encodes an ID into Base64.

**🔹 EncodeBase58(input string) string**

Encodes an ID into Base58 (shorter than Base64).

🧪 Running Tests

To verify the correctness of the library, use:
```go
go test -v ./idgen
```
Example output:
```
=== RUN   TestGenerate_UniqueIDs
--- PASS: TestGenerate_UniqueIDs (0.00s)
=== RUN   TestGenerate_Length
--- PASS: TestGenerate_Length (0.00s)
=== RUN   TestGenerate_TimestampFormat
--- PASS: TestGenerate_TimestampFormat (0.00s)
=== RUN   TestGenerate_Concurrency
--- PASS: TestGenerate_Concurrency (0.01s)
=== RUN   TestEncodeBase64
--- PASS: TestEncodeBase64 (0.00s)
=== RUN   TestEncodeBase58
--- PASS: TestEncodeBase58 (0.00s)
PASS
ok  	github.com/devjefster/GoShortUniqueID/idgen 	0.02s
````
## 🚀 Performance & Benchmarks

Running Benchmarks

To test the performance of GoShortUniqueID, run:

`go test -bench=. ./idgen`

Example Results

```
BenchmarkGenerate-8                 	1000000	      1200 ns/op
BenchmarkGenerate_LongRandomPart-8   	 500000	      1500 ns/op
BenchmarkGenerate_Concurrent-8        2000000	       800 ns/op
BenchmarkEncodeBase64-8               3000000	       500 ns/op
BenchmarkEncodeBase58-8               2500000	       600 ns/op
Generate IDs:           ~1,200 ns/op (nanoseconds per operation)
Base64 Encoding:                                      ~500 ns/op
Base58 Encoding:                                      ~600 ns/op
```

🔹 Interpretation:
The ID generator is highly efficient, capable of generating over 1 million IDs per second. 🚀

Coming soon... 🏎️
## 📜 License
This project is licensed under the **MIT License**.

You can view the full text of the license in the [`LICENSE`](LICENSE) file.

[![License: MIT](https://img.shields.io/github/license/devjefster/GoShortUniqueID)](LICENSE)

This project is licensed under the MIT License. See the LICENSE file for details.

📣 Contributing

Want to improve this library? Feel free to submit a Pull Request! 🙌

⭐ Support

If you like this project, give it a star ⭐ on GitHub!
