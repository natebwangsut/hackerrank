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

# stuart = non VOWEL
# kevin = VOWEL

VOWEL = "AEIOU"


def minion_game(s):
    stuart = 0
    kevin = 0

    for i in range(len(s)):
        if s[i] not in VOWEL:
            stuart += len(s) - i
        else:
            kevin += len(s) - i

    if stuart > kevin:
        print(f"Stuart {stuart}")
    elif stuart == kevin:
        print(f"Draw")
    else:
        print(f"Kevin {kevin}")


# Main
if __name__ == "__main__":
    line = input()
    print(minion_game(line))
