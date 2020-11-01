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


def custom_sort(x: str):
    """
    All sorted lowercase letters are ahead of uppercase letters.
    All sorted uppercase letters are ahead of digits.
    All sorted odd digits are ahead of sorted even digits.
    """
    return (
        not str(x).isalpha(),
        not str(x).islower(),
        str(x).isdigit(),
        int(x) % 2 == 0 if (str(x).isdigit()) else str(x),
        str(x),
    )


print(*sorted(input(), key=custom_sort), sep="")
