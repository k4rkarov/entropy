<h3 align="center">Entropy</h3>
<h1 align="center"> <img src="https://github.com/k4rkarov/entropy/blob/main/carbon.png" alt="procontor" width="400px"></h1>

A Go-based program crafted to evaluate the strength of passwords through the analysis of entropy and semantic robustness.

<br>

# Installation Instructions

`entropy` requires **Go 1.18** or later to install successfully. Run the following command to install the latest version: 

```sh
go install github.com/k4rkarov/entropy/cmd/entropy@latest
```

The `-s` flag relies on the `zxcvbn-go` library for semantic strength evaluation. The library will be automatically installed when you run `entropy` if it's not already available. Ensure your Go environment is properly configured to fetch and install dependencies.

# Usage

```sh
entropy -h
```

This will display the help menu.

```console

 ______       _                         
|  ____|     | |                        
| |__   _ __ | |_ _ __ ___  _ __  _   _ 
|  __| | '_ \| __| '__/ _ \| '_ \| | | |
| |____| | | | |_| | | (_) | |_) | |_| |
|______|_| |_|\__|_|  \___/| .__/ \__, |
                           | |     __/ |
                           |_|    |___/ 
 
       by k4rkarov (v2.0)

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

```

# Running entropy

### Calculate Password Entropy

```
$ entropy -p mypassword
Entropy: 47.00 bits
```

From a list

```
$ entropy -p -L list.txt 
password123 - Entropy: 56.87 bits
qwerty - Entropy: 28.20 bits
aloha@2024 - Entropy: 55.24 bits
impossible for you to know - Entropy: 123.63 bits
```

Adding verbosity

```
$ entropy -p mypassword -v
Entropy: 47.00 bits
Charset Size: 26
Length: 10

Lower Case Latin Alphabet (a-z)


$ entropy -p -L list.txt -v
password123 - Entropy: 56.87 bits
Charset Size: 36
Length: 11

Lower Case Latin Alphabet (a-z)
Numbers (0-9)

qwerty - Entropy: 28.20 bits
Charset Size: 26
Length: 6

Lower Case Latin Alphabet (a-z)

aloha@2024 - Entropy: 55.24 bits
Charset Size: 46
Length: 10

Lower Case Latin Alphabet (a-z)
Numbers (0-9)
Symbols (!@#$%()^&*)

impossible for you to know - Entropy: 123.63 bits
Charset Size: 27
Length: 26

Lower Case Latin Alphabet (a-z)
Space (' ')
```

### Calculate Entropy Based on Data Criteria

```
$ entropy -pc 14 lc uc d
Entropy: 83.36 bits
```

### Calculate Password's Semantic Strength

The `-s` option evaluates the semantic strength of the password using the `zxcvbn-go` library. Additionally, you can customize the list of common words by editing the `common_words.txt` file. You can add words from other languages or modify the list as needed. Ensure these modifications are made **before compiling the Go program**.

```
$ entropy -s Pass@123
Semantically weak.

$ entropy -s 'Ash&r$%D6D!@#18723'
PROBABLY NOT semantically weak
```

From a list

```
$ entropy -s -L list.txt 
password123 - Semantically weak.
qwerty - Semantically weak.
aloha@2024 - Semantically weak.
impossible for you to know - PROBABLY NOT semantically weak
```

Adding verbosity

```
$ entropy -s Pass@123 -v
Pass@123 - Semantically weak
Password is semantically weak according to zxcvbn.
Password contains a numeric sequence.
Password has a common word.

$ entropy -s -L list.txt -v
password123 - Semantically weak
Password is semantically weak according to zxcvbn.
Password contains a numeric sequence.
Password has a common word.
qwerty - Semantically weak
Password is semantically weak according to zxcvbn.
Password has a common word.
aloha@2024 - Semantically weak
Password is semantically weak according to zxcvbn.
Password has a common word.
Password has a year pattern.
impossible for you to know - PROBABLY NOT semantically weak
```

### Notes

- If your password contains any special characters, enclose the entire password in single quotes to avoid issues with the program. Special characters include: `!@#$%^&*(), \~-_=+[{]}|;:,<.>/?,` and spaces.

For any issues or feature requests, please open an issue on the [GitHub repository](https://github.com/k4rkarov/entropy).

Enjoy using Entropy to secure your passwords!