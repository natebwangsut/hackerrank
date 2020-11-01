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


def mat(height, width):

    midHeight = (height - 1) / 2
    midWidth = (width - 1) / 2

    for i in range(height):
        for j in range(width):

            # Welcome
            if i == (height - 1) / 2:
                extension = "".join(["-" for i in range(int((width - 15) / 2))])
                print(extension + "----WELCOME----" + extension, end="")
                break

            # Within Diamond
            if midWidth - abs(j - midWidth) >= abs(i - midHeight) * 3:
                if abs(j - midWidth) % 3 == 0:
                    print("|", end="")
                else:
                    print(".", end="")
            else:
                print("-", end="")

        # Print "\n" char
        print()


if __name__ == "__main__":
    inputString = input()
    height, width = inputString.split()
    mat(int(height), int(width))
