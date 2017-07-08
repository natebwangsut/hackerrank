/*
 * Nate B. Wangsutthitham
 * <@natebwangsut | nate.bwangsut@gmail.com>
 */

package main

import "fmt"

func gradingStudents() {

	var n, num int
	var grades []int

	// Input
	fmt.Scanf("%v", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &num)
		grades = append(grades, num)
	}

	// Output
	for i := 0; i < n; i++ {
		fmt.Println(roundGrade(grades[i]))
	}
}

func roundGrade(grade int) int {
	// Early exit if lower than 38
	if grade < 38 {
		return grade
	} else if (grade % 5) >= 3 {
		for (grade % 5) != 0 {
			grade++
		}
	}
	return grade
}
