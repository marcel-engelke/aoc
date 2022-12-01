import sys

def main():
	# 1
	elves = []
	elf = 0
	for line in sys.stdin.readlines():
		if (line == "\n"):
			elves.append(elf)
			elf = 0
			continue
		elf += int(line.replace("\n", ""))
	# print(max(elves))
	# 2
	cals = 0
	for _ in range(0, 3):
		cals += elves.pop(elves.index(max(elves)))
	print(cals)

main()