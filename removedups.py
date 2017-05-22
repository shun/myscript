#!/usr/local/bin/python3

import os

lines = []

f = open("$HOME/.bash_history")
line = f.readline()
while line:
    lines.append(line.strip())
    line = f.readline()

f.close()

# remove duplicated command
lines = sorted(set(lines), key=lines.index)

f = open(os.environ["HOME"] + "/a.txt", "w")
for l in lines:
    f.writelines(l + "\n")

f.close()

