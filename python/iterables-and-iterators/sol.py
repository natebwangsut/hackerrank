#!/usr/bin/env python

import itertools

line = input()
line = input().split()
repeat = int(input())

com = list(itertools.combinations(line, repeat))

print(sum(('a' in x) for x in com)/len(com))