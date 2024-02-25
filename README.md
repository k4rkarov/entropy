<h3 align="center">Entropy</h3>
<h1 align="center"> <img src="https://github.com/k4rkarov/entropy/blob/main/carbon.png" alt="procontor" width="400px"></h1>

A Go program designed to assess password strength by calculating entropy and semantic strength. This project is a work in progress, and contributions are welcome through pull requests :) 

<br>

# Installation Instructions

`entropy` requires **go1.18** to install successfully. Run the following command to install the latest version: 

```sh
go install github.com/k4rkarov/entropy/cmd/entropy@latest
```

Now you can move the binary to your /usr/bin folder so you can use it freely on your OS:

```
mv go/bin/entropy /usr/bin
```

# Usage

```sh
entropy -h
```

This will display help menu.

```console

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

Criteria (for -pe option):
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
  entropy -pc 14 lc uc d -v
  entropy -s Pass@123 -v
  entropy -s -L passwords.txt

```

# Running entropy

Calculate Password Entropy:

```
$ entropy -p mypassword
Entropy: 47.00 bits

$ entropy -p mypassword -v
Entropy: 47.00 bits
Charset Size: 26
Length: 10

Lower Case Latin Alphabet (a-z)
```

Calculate Entropy based on data criteria:

```
$ entropy -pc 14 lc uc d
Entropy: 83.36 bits

```

Calculate password's semantic strength:

```
$ entropy -s Pass@123
Semantically weak.

$ entropy -s Pass@123 -v
Semantically weak:
Password contains a numeric sequence. 
Password has a common word. 

$ entropy -s 'Ash&r$%D6D!@#18723'
PROBABLY NOT semantically weak, but needs further check

```
