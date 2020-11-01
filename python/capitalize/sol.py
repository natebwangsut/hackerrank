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

import math
import os
import random
import re
import sys

# Complete the solve function below.
def solve(s):
    return " ".join([i[0].upper() + i[1:] if len(i) > 0 else i for i in s.split(" ")])


# Main
if __name__ == "__main__":
    line = input()
    print(solve(line))
