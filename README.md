<h3 align="center">Entropy</h3>
<h1 align="center"> <img src="https://github.com/k4rkarov/entropy/blob/main/carbon.png" alt="procontor" width="400px"></h1>

A Go program designed to assess password strength by calculating entropy and semantic strength. This project is a work in progress, and contributions are welcome through pull requests :) 

<br>

# Download and install

First, clone the repository to your machine:

```
git clone https://github.com/k4rkarov/entropy.git
```

Then, build the go binary:

```
cd entropy/
go build -o entropy
```

Now you can move the binary to your /usr/bin folder so you can use it freely on your OS:

```
mv entropy /usr/bin
```

# Usage
```
$ entropy

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
```

# Example

1 Calculate Password Entropy:

```
$ entropy 1 mypassword
Entropy: 47.00 bits

$ entropy 1 mypassword -v
Entropy: 47.00 bits
Charset Size: 26
Length: 10

Lower Case Latin Alphabet (a-z)
```

2 Calculate Entropy based on data criteria:

```
$ entropy 2 14 lc uc d
Entropy: 83.36 bits

```

3 Calculate password's semantic strength:

```
$ entropy 3 Pass@123
Semantically weak: 
Password contains a numeric sequence. 
Password has a common word. 

$ entropy 3 'Ash&r$%D6D!@#18723'
PROBABLY NOT semantically weak, but needs further check
```
