#!/usr/bin/env python3
import re
import sys


# Part 1
def get_calibration_digits(line: str) -> list[int]: 
	nums = []
	for i in range(len(line)):
		try:
			num = int(line[i])
			nums.append(num)
			break
		except ValueError:
			continue
	for i in range(len(line) - 1, -1, -1):
		try:
			num = int(line[i])
			nums.append(num)
			break
		except ValueError:
			continue
	return nums
	
EXP = re.compile("one|two|three|four|five|six|seven|eight|nine|[0-9]")
WORDS = {
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
}

def get_calibration_number(line: str) -> int:
	m = re.findall(EXP, line)
	try:
		n1 = int(m[0])
	except ValueError:
		n1 = WORDS[m[0]]
		print("regex", n1)
	try:
		n2 = int(m[-1])
	except ValueError:
		n2 = WORDS[m[-1]]
		print("regex", n2)
	return int(f"{n1}{n2}")	

sum = 0
while True:
	try:
		line = sys.stdin.readline().strip()
		if len(line) == 0:
			break
		num = get_calibration_number(line)
		print(num)
		sum += num
	except EOFError:
		print("eof")
		break
print(sum)
	