#!/usr/bin/python3
import re
import sys

def read_line_cargo() -> list[int] | None:
    while True:
        cargo = []
        line = sys.stdin.readline()
        if line[1] == "1": break
        line = line.strip("\n")

        for i in range(1, 34, 4):
            cargo.append(line[i] != " " and line[i] or -1)
        yield cargo
    
def read_line_instr() -> list[int] | None:
    while True:
        instr = []
        line = sys.stdin.readline()
        if len(line) == 1 and line[0] == "\n": continue
        if not line: break
        for n in re.findall("\d+", line):
            instr.append(int(n))
        yield instr

def one() -> str:
    s = []
    for i in range(0, 9):
        s.append([])

    for c in read_line_cargo():
        for i in range(0, 9):
            if c[i] != -1:
                s[i].insert(0, c[i])
    
    for instr in read_line_instr():
        for i in range(0, instr[0]):
            s[instr[2] - 1].append(s[instr[1] - 1].pop())
    
    msg = ""
    for i in range(0, 9):
        msg += s[i].pop()
    return msg

def two() -> str:
    s = []
    for i in range(0, 9):
        s.append([])

    for c in read_line_cargo():
        for i in range(0, 9):
            if c[i] != -1:
                s[i].insert(0, c[i])
    
    for instr in read_line_instr():
        tmp = []
        for i in range(0, instr[0]):
            tmp.append(s[instr[1] - 1].pop())
        while(len(tmp) > 0):
            s[instr[2] - 1].append(tmp.pop())
    
    msg = ""
    for i in range(0, 9):
        msg += s[i].pop()
    return msg

print(two())