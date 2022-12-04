import sys

def read_line() -> list[int] | None:
	while True:
		line = sys.stdin.readline()
		if not line: break
		line = line.strip("\n ")
		pairs = line.split(",")
		elves = []
		for e in pairs[0].split("-"):
			elves.append(int(e))
		for e in pairs[1].split("-"):
			elves.append(int(e))
		yield elves

def one() -> int:
	contained = 0
	for asmt in read_line():
		if asmt[0] <= asmt[2] and asmt[1] >= asmt[3]:
			contained += 1
		elif asmt[2] <= asmt[0] and asmt[3] >= asmt[1]:
			contained += 1
	return contained

def two() -> int:
	overlapped = 0
	for asmt in read_line():
		if asmt[0] <= asmt[2] and asmt[1] >= asmt[3]:
			overlapped += 1
		elif asmt[2] <= asmt[0] and asmt[3] >= asmt[1]:
			overlapped += 1
		elif asmt[0] >= asmt[2] and asmt[0] <= asmt[3]:
			overlapped += 1
		elif asmt[1] >= asmt[2] and asmt[1] <= asmt[3]:
			overlapped += 1
		elif asmt[2] >= asmt[0] and asmt[2] <= asmt[1]:
			overlapped += 1
		elif asmt[3] >= asmt[0] and asmt[3] <= asmt[1]:
			overlapped += 1
			
	return overlapped

print(two())