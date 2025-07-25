package main

import (
	"bytes"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/argon2"
)

func hashPass(salt []byte, plainPassword string) []byte {
	hashedPass := argon2.IDKey([]byte(plainPassword), []byte(salt), 1, 64*1024, 4, 32)
	return append(salt, hashedPass...)
	// [salt] + [pass_hash]
}

func checkPass(passHash []byte, plainPassword string) bool {
	salt := make([]byte, 8)
	copy(salt, passHash[:8])
	userPassHash := hashPass(salt, plainPassword)
	return bytes.Equal(userPassHash, passHash)
}

func passExample() {
	pass := "love"

	salt := make([]byte, 8)
	rand.Read(salt)
	fmt.Printf("salt: %x\n\n", salt)

	hashedPass := hashPass(salt, pass)
	fmt.Printf("hashedPass: %x\n\n", hashedPass)

	passValid := checkPass(hashedPass, pass)
	fmt.Printf("OK passValid: %v\n\n", passValid)

	passValid = checkPass(hashedPass, "nolove")
	fmt.Printf("BAD passValid: %v\n\n", passValid)
}

func main2() {
	for i := 0; i < 3; i++ {
		fmt.Println("\titeration", i)
		passExample()
	}
}
