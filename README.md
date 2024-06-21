# pwdgen
 Golang CLI random password generator tool

The following flags are supported:
  -a, --alphanumeric   Use only Alphanumeric charaters (no special symbols)
  -c, --count int      Number of passwords to generate (default 1)
  -h, --help           help for pwdgen
  -l, --length int     Set the length of password(s) (default 16)


## todo

**config management**
create cli for changing the parameters in the config file, as well as reading them

**required patterns:**
set certain patterns that need to be met.
*these patterns can include:*
- at least 1 upper case letter
- at least 1 lower case letter
- at least 1 special character
- at least 1 number

**anti patterns:**
set certain patterns, so that the password gets rerolled if a match is found.
This is in order to avoid accidental rejection by overly complex requirements set in certain systems
Example: Pattern = three characters next to each other on the keyboard | Matched Password = 'Ower<dIqMX'
*these patterns can include:*
- numbers in increasing, decreasing order
- 3 consecutive characters from the abc
- the same 3 or more characters repeated after the other (asdasd)