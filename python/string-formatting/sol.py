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


def print_formatted(number):
    # your code goes here
    for i in range(1, number + 1):
        print(
            "{num:{fill}g} {num:{fill}o} {num:{fill}X} {num:{fill}b}".format(
                num=i, fill=len("{0:b}".format(number))
            )
        )


if __name__ == "__main__":
    n = int(input())
    print_formatted(n)
