/*
 * Nate B. Wangsutthitham
 * <@natebwangsut | nate.bwangsut@gmail.com>
 */

package main

import "fmt"

func sum(a uint32, b uint32) uint32 {
	// Hint: Type return (a+b) below
	return (a + b)
}

func main() {
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a, &b)
	res = sum(a, b)
	fmt.Println(res)
}
