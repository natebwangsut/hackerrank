package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the matrixRotation function below.
func matrixRotation(matrix [][]int32, r int32) {

	// Layers = how many rings to rotate (min of width or height / 2)
	h, w := len(matrix), len(matrix[0])
	layers := int(math.Min(float64(h/2), float64(w/2)))

	rings := [][]int32{}

	// For each layer
	for l := layers - 1; l >= 0; l-- {
		ring := []int32{}
		Lring := []int32{}
		Rring := []int32{}
		for i := l; i < h-l; i++ {
			// top
			if i == l {
				for j := l; j < w-l; j++ {
					ring = append(ring, matrix[i][j])
				}
				continue
			}
			// bottom
			if i == h-l-1 {
				ring = append(ring, Rring...)
				for j := w - l - 1; j >= l; j-- {
					ring = append(ring, matrix[i][j])
				}
				ring = append(ring, Lring...)
				continue
			}
			// middleSection
			Lring = append([]int32{matrix[i][l]}, Lring...)
			Rring = append(Rring, matrix[i][w-l-1])
		}
		rings = append(rings, rotateArr(ring, r)) // rotate each ring and store them in "rings"
	}

	// Now do assignment - for each layer
	for l := layers - 1; l >= 0; l-- {
		side := 1
		count := 0
		inverseL := layers - 1 - l
		for i := l; i < h-l; i++ {
			// top
			if i == l {
				for j := l; j < w-l; j++ {
					matrix[i][j] = rings[inverseL][count]
					count++
				}
				continue
			}
			// bottom
			if i == h-l-1 {
				for j := w - l - 1; j >= l; j-- {
					matrix[i][j] = rings[inverseL][count]
					count++
				}
				continue
			}
			// middleSection
			matrix[i][w-l-1] = rings[inverseL][count]
			count++
			matrix[i][l] = rings[inverseL][len(rings[inverseL])-side]
			side++
		}
	}
	printAns(matrix)
}

func rotateArr(arr []int32, r int32) []int32 {
	if len(arr) == 0 {
		return arr
	}
	rotation := int(r) % len(arr)
	return append(arr[rotation:], arr[:rotation]...)
}

func printAns(matrix [][]int32) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	mnr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	mTemp, err := strconv.ParseInt(mnr[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(mnr[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	rTemp, err := strconv.ParseInt(mnr[2], 10, 64)
	checkError(err)
	r := int32(rTemp)

	var matrix [][]int32
	for i := 0; i < int(m); i++ {
		matrixRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var matrixRow []int32
		for _, matrixRowItem := range matrixRowTemp {
			matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
			checkError(err)
			matrixItem := int32(matrixItemTemp)
			matrixRow = append(matrixRow, matrixItem)
		}

		if len(matrixRow) != int(n) {
			panic("Bad input")
		}

		matrix = append(matrix, matrixRow)
	}

	matrixRotation(matrix, r)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
