package main

import (
	"bytes"
	r "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)

// version history
// v0.1.0, 2023-01-10; initial release
// v0.1.1, 2023-01-11; added sanity checks to passwordLength & passwordCount inputs
// v0.1.2, 2023-02-09; use "crypto/rand" to create password seed

// clear screen function
func clearScreen() {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func generateSeedKey() int64 {
	b := make([]byte, 16)
	_, err := r.Read(b)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	var seed int64
	binary.Read(bytes.NewReader(b), binary.BigEndian, &seed)
	return seed
}

func RandPassGen(n int) string {
	var alphaChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var digitChars = []rune("0123456789")
	var specialChars = []rune("!@#$%^&*()_+-=[]{}\\|;':\"<>,.?/")

	b := make([]rune, n)
	for i := range b {
		seed := generateSeedKey()
		rand.Seed(seed)
		r := rand.Intn(3)
		if r == 0 {
			b[i] = alphaChars[rand.Intn(len(alphaChars))]
		} else if r == 1 {
			b[i] = digitChars[rand.Intn(len(digitChars))]
		} else {
			b[i] = specialChars[rand.Intn(len(specialChars))]
		}
	}
	return string(b)
}

func main() {
	var passwordLength int
	var passwordCount int
	clearScreen()
	fmt.Println(" ----------------- ")
	fmt.Println("| Cyclone PassGen |")
	fmt.Println(" ----------------- ")
	fmt.Println()
	for {
		fmt.Print("Password length: ")
		fmt.Scan(&passwordLength)
		if passwordLength < 8 || passwordLength > 1000 {
			fmt.Println("Invalid password length")
		} else {
			break
		}
	}

	for {
		fmt.Print("Number of passwords to generate: ")
		fmt.Scan(&passwordCount)
		if passwordCount < 1 || passwordCount > 100000 {
			fmt.Println("Invalid password count")
		} else {
			break
		}
	}

	for i := 0; i < passwordCount; i++ {
		fmt.Println(RandPassGen(passwordLength))
	}
	fmt.Println()
}
