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
  entropy <options> '<password length>' [criteria]
 
Options:
  1 - Calculate Password Entropy
  2 - Calculate Entropy based on data criteria
  3 - Calculate password's semantic strength
 
Criteria (optional):
  lc - lowercase characters
  uc - uppercase characters
  d - digits
  s - special characters = !@#$%^&*()
  sp - additional special characters = `~-_=+[{]}\|;:'",<.>/?
  spc - space
 
Example:
  entropy 1 mypassword
  entropy 1 'Pass@2#@!'
  entropy 2 14 lc uc d
  entropy 3 Pass@123
```

# Example

1 - Calculate Password Entropy:

```
$ entropy 1 mypassword
Your password contains:

Entropy: 47.00 bits
Charset Size: 26
Length: 10

Lower Case Latin Alphabet (a-z)
```

2 - Calculate Entropy based on data criteria:

```
$ entropy 2 14 lc uc d

Entropy: 83.36 bits

```

3 - Calculate password's semantic strength:

```
$ entropy 3 Pass@123
Password is semantically weak!

$ entropy 3 'Ash&r$%D6D!@#18723'
Password is PROBABLY NOT semantically weak!
```
