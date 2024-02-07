package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isNumericSequence(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i] >= '0' && password[i] <= '9' && password[i] == password[i+1]-1 && password[i] == password[i+2]-2 {
			return true
		}
	}
	return false
}

func isCommonWord(password string) bool {
	commonWords := []string{"senha", "password", "pass", "mudar", "change", "teste", "test"}
	for _, word := range commonWords {
		if strings.Contains(strings.ToLower(password), word) {
			return true
		}
	}
	return false
}

func isYearPattern(password string) bool {
	return regexp.MustCompile(`(?i)[a-z]+\@\d{4}`).MatchString(password)
}

func calculateSemanticStrength(password string) string {
	if isNumericSequence(password) || isCommonWord(password) || isYearPattern(password) {
		return "Password is semantically weak!"
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
 
       by k4rkarov (v1.0)


Usage:
  entropy <option> <password> [criteria] [-v]

Options:
  1       Calculate Password Entropy
  2       Calculate Entropy based on specified criteria
  3       Evaluate password's semantic strength

Criteria (for option 2):
  length  The number of characters in the password
  lc      lowercase characters: (a-z)
  uc      uppercase characters: (A-Z)
  d       digits: (0-9)
  s       special characters: !@#$%^&*()
  sp      additional special characters: ~-_=+[{]}|;:'",<.>/?
  spc     space (' ')

Output:
  -v      Increase verbosity level

Examples:
  entropy 1 mypassword
  entropy 1 'Pass@2#@!'
  entropy 2 14 lc uc d
  entropy 3 Pass@123

`)
}


func calculatePasswdEntropy(password string, verbose bool) string {
    length := len(password)
    passCharac := make([]string, 0)
    text := ""

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

    if verbose {
            text += fmt.Sprintf("Entropy: %.2f bits\nCharset Size: %d\nLength: %d\n\n", math.Round(math.Log2(math.Pow(float64(chars), float64(length)))*100)/100, chars, length)
    } else {
        text += fmt.Sprintf("Entropy: %.2f bits", math.Round(math.Log2(math.Pow(float64(chars), float64(length)))*100)/100)
    }

    if verbose {
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

	if length > 0 && (lowercase || uppercase || digit || special || specialPlus || space) {
    text += fmt.Sprintf("Entropy: %.2f bits", math.Round(math.Log2(math.Pow(float64(chars), float64(length)))*100)/100)
} else if length > 0 {
    text += "No valid criteria selected. Please select at least one criteria:\n" +
  "  lc  - lowercase characters: (a-z)\n" +
  "  uc  - uppercase characters: (A-Z)\n" +
  "  d   - digits: (0-9)\n" +
  "  s   - special characters: !@#$%^&*()\n" +
  "  sp  - additional special characters: ~-_=+[{]}|;:'\",<.>/?\n" +
  "  spc - space (' ')\n"

} else {
    text += "Missing password length for options 2."
}

	return text
}

func main() {
    if len(os.Args) < 2 {
        printHelp()
        os.Exit(1)
    }

    option := os.Args[1]
    verbose := false

    for i := 2; i < len(os.Args); i++ {
        if os.Args[i] == "-v" {
            verbose = true
            os.Args = append(os.Args[:i], os.Args[i+1:]...)
            i-- // Adjust the index since we removed an element
        }
    }

    switch option {
    case "1":
        if len(os.Args) < 3 {
            fmt.Println("Missing password for option 1.")
            os.Exit(1)
        }
        password := os.Args[2]
        result := calculatePasswdEntropy(password, verbose)
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

