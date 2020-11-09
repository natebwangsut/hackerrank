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

# Complete the matrixRotation function below.
def matrixRotation(matrix, r):

    h = len(matrix)
    w = len(matrix[0])
    layer = int(min([h / 2, w / 2]))
    rings = {}

    for l in range(0, layer):
        topRing = [matrix[l][j] for j in range(l, w - l)]
        bottomRing = [matrix[h - 1 - l][j] for j in range(w - 1 - l, l - 1, -1)]

        leftRing = [matrix[i][l] for i in range(l + 1, h - 1 - l)]
        leftRing.reverse()
        rightRing = [matrix[i][w - 1 - l] for i in range(l + 1, h - 1 - l)]

        fullRing = [*topRing, *rightRing, *bottomRing, *leftRing]
        actualRotation = r % len(fullRing)
        rings[l] = [*fullRing[actualRotation:], *fullRing[:actualRotation]]

    for l in range(0, layer):
        side = 1
        count = 0
        for i in range(l, h - l):

            if i == l:
                for j in range(l, w - l):
                    matrix[i][j] = rings[l][count]
                    count += 1
            elif i == h - l - 1:
                for j in range(w - l - 1, l - 1, -1):
                    matrix[i][j] = rings[l][count]
                    count += 1
            else:
                matrix[i][w - l - 1] = rings[l][count]
                count += 1
                matrix[i][l] = rings[l][len(rings[l]) - side]
                side += 1

    for i in range(0, len(matrix)):
        for j in range(0, len(matrix[i])):
            print(matrix[i][j], end=" ")
        print()


# Main
if __name__ == "__main__":
    mnr = input().rstrip().split()

    m = int(mnr[0])
    n = int(mnr[1])
    r = int(mnr[2])

    matrix = []
    for _ in range(m):
        matrix.append(list(map(int, input().rstrip().split())))

    matrixRotation(matrix, r)
