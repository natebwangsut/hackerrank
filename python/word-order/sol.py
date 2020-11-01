#!/usr/bin/env python

"""
Hackerrank Solution
-
Nate B. Wangsutthitham
<@natebwangsut | nate.bwangsut@gmail.com>
"""

__author__ = "Nate B. Wangsutthitham"
__credits__ = ["Nate B. Wangsutthitham"]
__license__ = "MIT"
__version__ = "1.0.0"
__maintainer__ = "Nate B. Wangsutthitham"
__email__ = "nate.bwangsut@gmail.com"

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
