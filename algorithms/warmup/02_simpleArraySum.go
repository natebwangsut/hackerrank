/*
 * Nate B. Wangsutthitham
 * <@natebwangsut | nate.bwangsut@gmail.com>
 */

package main

import "fmt"

func simpleArraySum() {
	//enter your code here. Read input from STDIN. Print output to STDOUT
	var num, nextnum int
	fmt.Scanf("%v\n", &num)

	sum := 0
	for i := 0; i < num; i++ {
		fmt.Scanf("%v", &nextnum)
		sum += nextnum
	}
	fmt.Println(sum)
}
