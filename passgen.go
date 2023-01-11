package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// version history
// v0.1.0, 2023-01-10; initial release
// v0.1.1, 2023-01-11; added sanity checks to passwordLength & passwordCount inputs

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

func RandPassGen(n int) string {
	var alphaChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var digitChars = []rune("0123456789")
	var specialChars = []rune("!@#$%^&*()_+-=[]{}\\|;':\"<>,.?/")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
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
