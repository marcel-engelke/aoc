import sys

symbols = {
	'X': 1,
	'Y': 2,
	'Z': 3
}

matches = {
	'A': {
		'X': 3,
		'Y': 6,
		'Z': 0
	},
	'B': {
		'X': 0,
		'Y': 3,
		'Z': 6
	},
	'C': {
		'X': 6,
		'Y': 0,
		'Z': 3
	}
}

def one():
	score = 0
	for line in sys.stdin.readlines():
		score += symbols[line[2]]
		score += matches[line[0]][line[2]]
	print(score)

# one()

outcomes = {
	'X': 0,
	'Y': 3,
	'Z': 6
}

matches = {
	'X': {
		'A': 3,
		'B': 1,
		'C': 2
	},
	'Y': {
		'A': 1,
		'B': 2,
		'C': 3
	},
	'Z': {
		'A': 2,
		'B': 3,
		'C': 1
	}
}

def two():
	score = 0
	for line in sys.stdin.readlines():
		score += outcomes[line[2]]
		score += matches[line[2]][line[0]]
	print(score)

two()
