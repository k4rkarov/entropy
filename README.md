# entropy
A Go program designed to calculate both the password entropy and its semantic strength.

# Install

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

# Main screen
```
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
