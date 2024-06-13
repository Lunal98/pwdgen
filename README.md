# pwdgen
 Golang CLI random password generator tool



## todo

config management:
	save the configured parameters into a yaml so it doesn't need to be typed every time
Basic Parameteres:
	length
	complexity
	number of passwords generated
required patterns:
	set certain patterns that need to be met.
	these patterns can include:
	- at least 1 upper case letter
	- at least 1 lower case letter
	- at least 1 special character
	- at least 1 number
anti patterns:
	set certain patterns, so that the password gets rerolled if a match is found.
	This is in order to avoid accidental rejection by overly complex requirements set in certain systems
	Example: Pattern = three characters next to each other on the keyboard | Matched Password = 'Ower<dIqMX'
	these patterns can include:
	- numbers in increasing, decreasing order
	- 3 consecutive characters from the abc
	- the same 3 or more characters repeated after the other (asdasd)