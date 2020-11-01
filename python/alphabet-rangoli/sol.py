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

def print_rangoli(size):

    h = size * 2 - 1
    w = size * 4 - 3

    midHeight = (h - 1) / 2
    midWidth = (w - 1) / 2

    for i in range(h):
        for j in range(w):
            if midWidth - abs(j - midWidth) >= abs(i - midHeight) * 2:
                if abs(j - midWidth) % 2 == 0:
                    print(
                        chr(int(abs(j - midWidth) / 2 + abs(i - midHeight)) + 97),
                        end="",
                    )
                else:
                    print("-", end="")
            else:
                print("-", end="")
        print()


if __name__ == "__main__":
    n = int(input())
    print_rangoli(n)
