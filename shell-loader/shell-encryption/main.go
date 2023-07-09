package main

import (
	"os"
)

const output = "output.sh"
const key = "myKey"

// xor runs a XOR encryption on the input string, encrypting it if it hasn't already been,
// and decrypting it if it has, using the key provided.
func xor(input, key string) (output string) {
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}

	return output
}

// Encrypt a shell script and output it using xor
func main() {
	if len(os.Args) != 2 {
		panic("This program expect a filename as first argument")
	}

	filename := os.Args[1]
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	encryptedContent := xor(string(file), key)

	if err := os.WriteFile(output, []byte(encryptedContent), 0666); err != nil {
		panic(err)
	}
}
