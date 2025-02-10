package main

import (
	"fmt"
	"github.com/devjefster/GoShortUniqueID/idgen"
)

func main() {
	// Create a generator with different settings
	idGen := idgen.New(6, "", "060102150405") // Default YYMMDDHHmmss format

	// Generate multiple unique IDs
	for i := 0; i < 5; i++ {
		id := idGen.Generate()
		fmt.Println("Raw ID:", id)
		fmt.Println("Base64:", idgen.EncodeBase64(id))
		fmt.Println("Base58:", idgen.EncodeBase58(id))
		fmt.Println("-----------")
	}
}
