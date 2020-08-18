/*
 * Nate B. Wangsutthitham
 * <@natebwangsut | nate.bwangsut@gmail.com>
 */

package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var num int
	fmt.Scanf("%v", &num)
	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			if num >= j+i {
				fmt.Printf("%s", " ")
			} else {
				fmt.Printf("%s", "#")
			}
		}
		fmt.Printf("%s", "\n")
	}
}
