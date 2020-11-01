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
