#!/usr/bin/python3
import sys
import typing

def read_line() -> str:
	line = sys.stdin.readline()
	return line.strip("\n")

def one():
	msg = read_line()
	chars = []
	i = 0
	for i in range(0, 4):
		chars.append(msg[i])
	for i in range(4, len(msg)):
		try:
			chars.index(msg[i])
		except ValueError:
			if is_unique(chars):
				return i
		chars.pop(0)
		chars.append(msg[i])
	
def two():
	msg = read_line()
	chars = []
	i = 0
	for i in range(0, 14):
		chars.append(msg[i])
	for i in range(14, len(msg)):
		try:
			chars.index(msg[i])
		except ValueError:
			if is_unique(chars):
				return i
		chars.pop(0)
		chars.append(msg[i])
		if is_unique(chars):
			return i + 1

def is_unique(chars: [chr]) -> bool:
	found = {}
	for c in chars:
		if found.get(c):
			return False
		found[c] = True
	return True

print(two())
