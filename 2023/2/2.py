#!/usr/bin/env python3

CUBES = {
	"red": 12,
	"green": 13,
	"blue": 14
}

# Part 1
def get_possible_id(line: str) -> int:
	line = line.split(":")
	rounds = line[1].strip().split(";")
	for r in rounds:
		r = r.strip()
		cubes = r.split(",")
		for c in cubes:
			c = c.strip()
			c = c.split(" ")
			if int(c[0]) > CUBES[c[1]]:
				return 0
	return int(line[0].split(" ")[1])
		

# Part 2
def get_game_power(line: str) -> int:
	line = line.split(":")
	min_cubes = {
		"red": 0,
		"green": 0,
		"blue": 0
	}
	rounds = line[1].strip().split(";")
	for r in rounds:
		r = r.strip()
		cubes = r.split(",")
		for c in cubes:
			c = c.strip()
			c = c.split(" ")
			min_cubes[c[1]] = max(min_cubes[c[1]], int(c[0]))
	return min_cubes["blue"] * min_cubes["green"] * min_cubes["red"]


lines = []
with open("input", "r") as f:
	lines = f.readlines()

sum_ids = 0
sum_powers = 0
for line in lines:
	game_id = get_possible_id(line)
	sum_ids += game_id

	game_power = get_game_power(line)
	sum_powers += game_power

print(sum_ids)
print(sum_powers)