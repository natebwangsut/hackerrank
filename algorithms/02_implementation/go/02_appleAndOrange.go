/*
 * Nate B. Wangsutthitham
 * <@natebwangsut | nate.bwangsut@gmail.com>
 */

package main

import "fmt"

func appleAndOrange() {

	var num int

	var s, t int
	fmt.Scanf("%v %v", &s, &t)

	var a, b int
	fmt.Scanf("%v %v", &a, &b)

	var m, n int
	fmt.Scanf("%v %v", &m, &n)

	var apple, orange int

	//var distanceFromA []int
	for i := 0; i < m; i++ {
		fmt.Scanf("%v", &num)
		distance := a + num
		if distance >= s && distance <= t {
			apple++
		}
		//distanceFromA = append(distanceFromA, num)
	}

	//var distanceFromB []int
	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &num)
		distance := b + num
		if distance >= s && distance <= t {
			orange++
		}
		//distanceFromB = append(distanceFromB, num)
	}

	fmt.Println(apple)
	fmt.Println(orange)
}
