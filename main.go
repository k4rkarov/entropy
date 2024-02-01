package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var semanticallyWeakPasswords = []string{"password123", "senha@2024", "123456"}

func calculatePasswdEntropy(password string) string {
	length := len(password)
	passCharac := make([]string, 0)
	text := "Your password contains:\n"

	var lowerCase, upperCase, numbers, specialChars, specialCharsPlus, space, others, chars int

	for i := 0; i < length; i++ {
		c := string(password[i])

		if lowerCase == 0 && strings.Contains("abcdefghijklmnopqrstuvwxyz", c) {
			chars += 26
			lowerCase = 1
		}
		if upperCase == 0 && strings.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZ", c) {
			chars += 26
			upperCase = 1
		}
		if numbers == 0 && strings.Contains("0123456789", c) {
			chars += 10
			numbers = 1
		}
		if specialChars == 0 && strings.Contains("!@#$%^&*()", c) {
			chars += 10
			specialChars = 1
		}
		if specialCharsPlus == 0 && strings.Contains("`~-_=+[{]}\\|;:'\",<.>/?", c) {
			chars += 22
			specialCharsPlus = 1
		}
		if space == 0 && c == " " {
			chars += 1
			space = 1
		}
		if others == 0 && (c < " " || c > "~") {
			chars += 32 + 128
			others = 1
		}
	}

	text += fmt.Sprintf("\nEntropy: %.2f bits\nCharset Size: %d\nLength: %d\n\n", math.Round(math.Log2(math.Pow(float64(chars), float64(length)))*100)/100, chars, length)

	if lowerCase > 0 {
		passCharac = append(passCharac, "Lower Case Latin Alphabet (a-z)")
	}
	if upperCase > 0 {
		passCharac = append(passCharac, "Upper Case Latin Alphabet (A-Z)")
	}
	if numbers > 0 {
		passCharac = append(passCharac, "Numbers (0-9)")
	}
	if specialChars > 0 {
		passCharac = append(passCharac, "Symbols (!@#$%()^&*)")
	}
	if specialCharsPlus > 0 {
		passCharac = append(passCharac, "Special Chars (`~-_=+[{]}\\|;:'\",<.>/?)")
	}
	if space > 0 {
		passCharac = append(passCharac, "Space (' ')")
	}
	if others > 0 {
		passCharac = append(passCharac, "Others")
	}

	for _, v := range passCharac {
		text += fmt.Sprintf("%s\n", v)
	}

	return text
}

func calculateEntropy(length int, lowercase, uppercase, digit, special, specialPlus, space bool) string {
	text := ""
	chars := 0

	if lowercase {
		chars += 26
	}
	if uppercase {
		chars += 26
	}
	if digit {
		chars += 10
	}
	if special {
		chars += 10
	}
	if specialPlus {
		chars += 22
	}
	if space {
		chars += 1
	}

	if length > 0 {
		text += fmt.Sprintf("\nEntropy: %.2f bits\n", math.Round(math.Log2(math.Pow(float64(chars), float64(length)))*100)/100)
	} else {
		text += "No criteria selected. Unable to calculate entropy.\n"
	}

	return text
}

func calculateSemanticStrength(password string) string {
	for _, weakPassword := range semanticallyWeakPasswords {
		if password == weakPassword {
			return "Password is semantically weak!"
		}
	}

	return "Password is PROBABLY NOT semantically weak!"
}

func printHelp() {
    fmt.Println(`
  ______       _                         
 |  ____|     | |                        
 | |__   _ __ | |_ _ __ ___  _ __  _   _ 
 |  __| | '_ \| __| '__/ _ \| '_ \| | | |
 | |____| | | | |_| | | (_) | |_) | |_| |
 |______|_| |_|\__|_|  \___/| .__/ \__, |
                            | |     __/ |
                            |_|    |___/ 
	`)
	fmt.Println("    by k4rkarov (v1.0)")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Usage:")
	fmt.Println("  entropy <options> '<password length>' [criteria]")
	fmt.Println(" ")
	fmt.Println("Options:")
	fmt.Println("  1 - Calculate Password Entropy")
	fmt.Println("  2 - Calculate Entropy based on data criteria")
	fmt.Println("  3 - Calculate password's semantic strength")
	fmt.Println(" ")
	fmt.Println("Criteria (optional):")
	fmt.Println("  lc - lowercase characters")
	fmt.Println("  uc - uppercase characters")
	fmt.Println("  d - digits")
	fmt.Println("  s - special characters = !@#$%^&*()")
    fmt.Println("  sp - additional special characters = `~-_=+[{]}\\|;:'\",<.>/?")
	fmt.Println("  spc - space")
	fmt.Println(" ")
	fmt.Println("Example:")
	fmt.Println("  entropy 1 mypassword")
	fmt.Println("  entropy 1 'Pass@2#@!'")
	fmt.Println("  entropy 2 14 lc uc d")
	fmt.Println("  entropy 3 Pass@123")
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	option := os.Args[1]

	switch option {
	case "1":
		if len(os.Args) < 3 {
			fmt.Println("Missing password for option 1.")
			os.Exit(1)
		}
		password := os.Args[2]
		result := calculatePasswdEntropy(password)
		fmt.Println(result)
	case "2":
		length := 0
		if len(os.Args) >= 3 {
			lengthStr := os.Args[2]
			val, err := strconv.Atoi(lengthStr)
			if err == nil && val > 0 {
				length = val
			}
		}

		lowercase := false
		uppercase := false
		digit := false
		special := false
		specialPlus := false
		space := false

		for i := 3; i < len(os.Args); i++ {
			switch os.Args[i] {
			case "lc":
				lowercase = true
			case "uc":
				uppercase = true
			case "d":
				digit = true
			case "s":
				special = true
			case "sp":
				specialPlus = true
			case "spc":
				space = true
			}
		}

		result := calculateEntropy(length, lowercase, uppercase, digit, special, specialPlus, space)
		fmt.Println(result)
	case "3":
		if len(os.Args) < 3 {
			fmt.Println("Missing password for option 3.")
			os.Exit(1)
		}
		password := os.Args[2]
		result := calculateSemanticStrength(password)
		fmt.Println(result)
	default:
		fmt.Println("Invalid option.")
	}
}

