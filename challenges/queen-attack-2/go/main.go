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

// Complete the queensAttack function below.
func queensAttack(n int32, k int32, rowQueen int32, colQueen int32, obstacles [][]int32) int32 {
	limits := cleanInput(n, k, rowQueen, colQueen, obstacles)
	count := int32(0)
	for _, v := range limits {
		count += v
	}
	return count
}

func cleanInput(n int32, k int32, rowQueen int32, colQueen int32, obstacles [][]int32) [8]int32 {
	limits := [8]int32{
		n - rowQueen,                // + 0
		rowQueen - 1,                // - 0
		n - colQueen,                // 0 +
		colQueen - 1,                // 0 -
		min(n-rowQueen, n-colQueen), // + +
		min(n-rowQueen, colQueen-1), // + -
		min(rowQueen-1, n-colQueen), // - +
		min(rowQueen-1, colQueen-1), // - -
	}

	//
	for _, values := range obstacles {

		rPos := values[0] - rowQueen - 1
		rNeg := rowQueen - values[0] - 1
		cPos := values[1] - colQueen - 1
		cNeg := colQueen - values[1] - 1

		// C0
		if colQueen == values[1] {
			// R+
			if limits[0] > rPos && values[0] > rowQueen {
				limits[0] = rPos
			}
			// R-
			if limits[1] > rNeg && values[0] < rowQueen {
				limits[1] = rNeg
			}
		}

		// R0
		if rowQueen == values[0] {
			// C+
			if limits[2] > cPos && values[1] > colQueen {
				limits[2] = cPos
			}
			// C-
			if limits[3] > cNeg && values[1] < colQueen {
				limits[3] = cNeg
			}
		}

		// Diagonal
		if intAbs(values[0]-rowQueen) == intAbs(values[1]-colQueen) {
			// R+ C+
			if limits[4] > rPos && values[0] > rowQueen &&
				limits[4] > cPos && values[1] > colQueen {
				limits[4] = min(rPos, cPos)
			}
			// R+ C-
			if limits[5] > rPos && values[0] > rowQueen &&
				limits[5] > cNeg && values[1] < colQueen {
				limits[5] = min(rPos, cNeg)
			}
			// R- C+
			if limits[6] > rNeg && values[0] < rowQueen &&
				limits[6] > cPos && values[1] > colQueen {
				limits[6] = min(rNeg, cPos)
			}
			// R- C-
			if limits[7] > rNeg && values[0] < rowQueen &&
				limits[7] > cNeg && values[1] < colQueen {
				limits[7] = min(rNeg, cNeg)
			}
		}
	}
	return limits
}

func min(v0, v1 int32) int32 {
	if v0 < v1 {
		return v0
	}
	return v1
}

func intAbs(num int32) float64 {
	return math.Abs(float64(num))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	r_qC_q := strings.Split(readLine(reader), " ")

	r_qTemp, err := strconv.ParseInt(r_qC_q[0], 10, 64)
	checkError(err)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(r_qC_q[1], 10, 64)
	checkError(err)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(readLine(reader), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
			checkError(err)
			obstaclesItem := int32(obstaclesItemTemp)
			obstaclesRow = append(obstaclesRow, obstaclesItem)
		}

		if len(obstaclesRow) != int(2) {
			panic("Bad input")
		}

		obstacles = append(obstacles, obstaclesRow)
	}

	result := queensAttack(n, k, r_q, c_q, obstacles)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
