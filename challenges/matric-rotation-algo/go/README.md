
# Matric Rotation

## Logics

```console
6 6 1
 1  2  3  4  5  6
 7  8  9 10 11 12
13 14 15 16 17 18
19 20 21 22 23 24
25 26 27 28 29 30
31 32 33 34 35 36
```

- For each layers starting from 3 (deepest) to 1 (shallowest)
put the items into a linked list in circular motion

```console
Layer 3: [15, 16, 21, 22]
Layer 2: [8, 9, 10, 11, 17, 23, 29, 28, 27, 26, 20, 14]
Layer 1: [1, 2, 3, 4, 5, 6, 12, 18, 24, 30, 36, 35, 34, 33, 32, 31, 25, 19, 13, 7]
```

- Rotate the list by using slicing properties
- Reuse the method to construct the list to replace the values to the original matrix

## Test Cases

```console
4 6 1
 1  2  3  4  5  6
 7  8  9 10 11 12
13 14 15 16 17 18
19 20 21 22 23 24
```

```console
4 4 1
1 2 3 4
5 6 7 8
9 10 11 12
13 14 15 16
```
