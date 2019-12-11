package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the queensAttack function below.
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	count := int32(0)
	count += verticalAttack(n, k, r_q, c_q, obstacles)
	count += horizontalAttack(n, k, r_q, c_q, obstacles)
	count += diagonalAttack(n, k, r_q, c_q, obstacles)
	return count
}

func verticalAttack(n int32, k int32, r int32, c int32, obstacles [][]int32) int32 {
	count := 0
Pos:
	for i := r + 1; i <= n; i++ {
		for _, value := range obstacles {
			// fmt.Print(value)
			if i == value[0] && c == value[1] {
				break Pos
			}
		}
		fmt.Print(i, c)
		count++
	}
Neg:
	for i := r - 1; i > 0; i-- {
		for _, value := range obstacles {
			// fmt.Print(value)
			if i == value[0] && c == value[1] {
				break Neg
			}
		}
		fmt.Print(i, c)
		count++
	}
	return int32(count)
}

func horizontalAttack(n int32, k int32, r int32, c int32, obstacles [][]int32) int32 {
	count := 0
Pos:
	for i := r + 1; i <= n; i++ {
		// fmt.Print(i, c)
		for _, value := range obstacles {
			// fmt.Print(value)
			if r == value[0] && i == value[1] {
				break Pos
			}
		}
		count++
	}
Neg:
	for i := r - 1; i > 0; i-- {
		// fmt.Print(i, c)
		for _, value := range obstacles {
			// fmt.Print(value)
			if r == value[0] && i == value[1] {
				break Neg
			}
		}
		count++
	}
	return int32(count)
}

func diagonalAttack(n int32, k int32, r int32, c int32, obstacles [][]int32) int32 {
	count := 0
PosPos:
	for i := int32(1); r+i <= n && c+i <= n; i++ {
		for _, value := range obstacles {
			// fmt.Print(value)
			if r+i == value[0] && c+i == value[1] {
				break PosPos
			}
		}
		fmt.Print(r+i, c+i)
		count++
	}
PosNeg:
	for i := int32(1); r+i <= n && 0 < c-i; i++ {
		for _, value := range obstacles {
			// fmt.Print(value)
			if r+i == value[0] && c-i == value[1] {
				break PosNeg
			}
		}
		fmt.Print(r+i, c-i)
		count++
	}
NegPos:
	for i := int32(1); 0 < r-i && c+i <= n; i++ {
		for _, value := range obstacles {
			// fmt.Print(value)
			if r-i == value[0] && c+i == value[1] {
				break NegPos
			}
		}
		fmt.Print(r-i, c+i)
		count++
	}
NegNeg:
	for i := int32(1); 0 < r-i && 0 < c-i; i++ {
		for _, value := range obstacles {
			// fmt.Print(value)
			if r-i == value[0] && c-i == value[1] {
				break NegNeg
			}
		}
		fmt.Print(r-i, c-i)
		count++
	}
	return int32(count)
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
