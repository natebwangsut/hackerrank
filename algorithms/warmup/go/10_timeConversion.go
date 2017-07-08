/*
 * Nate B. Wangsutthitham
 * <@natebwangsut | nate.bwangsut@gmail.com>
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var time string
	fmt.Scanf("%v", &time)

	slice := strings.Split(time, ":")

	hour, _ := strconv.Atoi(slice[0])
	minute := slice[1]
	second := slice[2][:2]

	if slice[2][2] == 'P' {
		if hour != 12 {
			hour += 12
			hour = hour % 24
		}
	} else {
		if hour == 12 {
			hour = 0
		}
	}

	fmt.Printf("%02v:%2v:%2v", hour, minute, second)
}
