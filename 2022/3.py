import sys

def gen_prios() -> dict:
	prios = {}
	for i in range(97,123):
		prios[chr(i)] = i - 96
	for i in range(65,91):
		prios[chr(i)] = i - 38
	return prios

def read_line() -> str | None:
	line = sys.stdin.readline()
	if not line:
		return
	yield line.strip(" \n")

def one() -> int:
	prio_sum = 0
	for line in read_line():
		prios = gen_prios()

		r1 = {}
		r2 = {}
		seen = {}
		for c in line[:int(len(line) / 2)]:
			r1[c] = True
		for c in line[int(len(line) / 2):]:
			r2[c] = True

		for c in r1:
			if r2.get(c) and not seen.get(c):
				prio_sum += prios[c]
				seen[c] = True
	return prio_sum

# print(one())

def two() -> int:
	prio_sum = 0
	prios = gen_prios()
	while True:
		line = read_line()
		if not line:
			break

		r1 = {}
		for c in line:
			r1[c] = True
		line = read_line()
		
		r2 = {}
		for c in line:
			r2[c] = True			
		line = read_line()

		r3 = {}
		for c in line:
			r3[c] = True			
		
		for c in r1:
			if r2.get(c) and r3.get(c):
				prio_sum += prios[c]

	return prio_sum

print(two())