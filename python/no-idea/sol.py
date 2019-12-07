#!/usr/bin/env python

# Forget about first line
line = input()

# Second line: n array
line = input()
n_list = line.split(" ")

# Third line: a set
line = input()
a_set = set(line.split(" "))

# Fourth line: b set
line = input()
b_set = set(line.split(" "))

print(sum((x in a_set) - (x in b_set) for x in n_list))