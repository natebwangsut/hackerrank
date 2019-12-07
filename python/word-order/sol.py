#!/usr/bin/env python

count = {}
lines = int(input())

for i in range(lines):
    word = input()
    try:
        count[word] += 1
    except KeyError:
        count[word] = 1

print(len(count.values()))
print(" ".join(map(str, count.values())))