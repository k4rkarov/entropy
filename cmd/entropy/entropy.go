package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	_ "embed"
)

//go:embed common_words.txt
var commonWordsContent string

// ANSI escape codes for colors
const (
	redColor   = "\033[1;31m"
	greenColor = "\033[1;32m"
	resetColor = "\033[0m"
)

func addRedColor(message string) string {
	return redColor + message + resetColor
}

func addGreenColor(message string) string {
	return greenColor + message + resetColor
}


func isCommonWord(password, filePath string) (bool, error) {
	// Use the embedded content
	words := strings.Fields(commonWordsContent)

	// Check if the password contains any common word
	for _, word := range words {
		if strings.Contains(strings.ToLower(password), strings.ToLower(word)) {
			return true, nil
		}
	}

	return false, nil
}

func isNumericSequence(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if (password[i] >= '0' && password[i] <= '9' &&
			(password[i]+1 == password[i+1] && password[i]+2 == password[i+2] ||
				password[i]-1 == password[i+1] && password[i]-2 == password[i+2])) {
			return true
		}
	}
	return false
}

func isYearPattern(password string) bool {
	return regexp.MustCompile(`(?i).*\d{4}`).MatchString(password)
}

func isAlphabeticSequence(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if (password[i] >= 'a' && password[i] <= 'z' &&
			(password[i]+1 == password[i+1] && password[i]+2 == password[i+2] ||
				password[i]-1 == password[i+1] && password[i]-2 == password[i+2])) ||
			(password[i] >= 'A' && password[i] <= 'Z' &&
				(password[i]+1 == password[i+1] && password[i]+2 == password[i+2] ||
					password[i]-1 == password[i+1] && password[i]-2 == password[i+2])) {
			return true
		}
	}
	return false
}

func calculateSemanticStrength(password string, verbose bool) string {
	var weaknessMessage string

	if isNumericSequence(password) {
		weaknessMessage += "\nPassword contains a numeric sequence."
	}

	commonWordFilePath := "common_words.txt"
	isCommon, err := isCommonWord(password, commonWordFilePath)
	if err != nil {
		return "Error reading common words file."
	}
	if isCommon {
		weaknessMessage += "\nPassword has a common word."
	}
	if isYearPattern(password) {
		weaknessMessage += "\nPassword has a year pattern."
	}
	if isAlphabeticSequence(password) {
		weaknessMessage += "\nPassword contains an alphabetic sequence."
	}

	if verbose && weaknessMessage != "" {
		return addRedColor("Semantically weak") + weaknessMessage
	} else if weaknessMessage != "" {
		return addRedColor("Semantically weak")
	}

	return addGreenColor("PROBABLY NOT semantically weak")
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
  entropy <option> <password> [criteria] [-L <file>] [-v]

Options:
  -p       Calculate Password Entropy
  -pc      Calculate Entropy based on specified criteria
  -s       Evaluate password's semantic strength

Criteria (for -pc option):
  length  The number of characters in the password
  lc      lowercase characters: (a-z)
  uc      uppercase characters: (A-Z)
  d       digits: (0-9)
  s       special characters: !@#$%^&*()
  sp      additional special characters: ~-_=+[{]}|;:'",<.>/?
  spc     space (' ')

Output:
  -v      Increase verbosity level
  -L      Specify a file with a list of passwords

Examples:
  entropy -p password123
  entropy -p 'Pass@2#@!' -v
  entropy -pc 14 lc uc d
  entropy -s Pass@123 -v
  entropy -s -L passwords.txt
`)
}


func readPasswordsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var passwords []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return passwords, nil
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
	passwordListFile := ""
	passwords := []string{}

	for i := 2; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-v":
			verbose = true
		case "-L": // Corrected the flag here
			if i+1 < len(os.Args) {
				passwordListFile = os.Args[i+1]
				i++ // Increment i to skip the next argument, which is the file path
			} else {
				fmt.Println("Missing file path after -L flag.")
				os.Exit(1)
			}
		default:
			passwords = append(passwords, os.Args[i])
		}
	}

	if passwordListFile != "" {
		filePasswords, err := readPasswordsFromFile(passwordListFile)
		if err != nil {
			fmt.Println("Error reading passwords from file:", err)
			os.Exit(1)
		}
		passwords = append(passwords, filePasswords...)
	}

	switch option {
	case "-p":
		if len(passwords) == 0 {
			printHelp()
			os.Exit(1)
		}
		for _, password := range passwords {
			result := calculatePasswdEntropy(password, verbose)
			if passwordListFile != "" {
				fmt.Printf("%s - %s\n", password, result)
			} else {
				fmt.Println(result)
			}
		}
	case "-pc":
		length := 0
		if len(os.Args) >= 3 {
			lengthStr := os.Args[2]
			val, err := strconv.Atoi(lengthStr)
			if err == nil && val > 0 {
				length = val
			}
		}

		if length == 0 {
			printHelp()
			os.Exit(1)
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
	case "-s":
		if len(passwords) == 0 {
			printHelp()
			os.Exit(1)
		}
		for _, password := range passwords {
			result := calculateSemanticStrength(password, verbose)
			if passwordListFile != "" {
				fmt.Printf("%s - %s\n", password, result)
			} else {
				fmt.Println(result)
			}
		}
	default:
		printHelp()
		os.Exit(1)
	}
}

