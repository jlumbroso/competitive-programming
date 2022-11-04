/* Write a program that reads three integers from standard
input and puts them into three variables A, B and C. */

// https://codeforces.com/problemset/problem/1624/A

/*

3
6
3 4 2 4 1 2
3
1000 1002 998
2
12 11

*/

package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	// write a loop to read the test cases
	for i := 0; i < t; i++ {
		// read the number of elements
		var n int
		fmt.Scan(&n)

		// read the elements of the array
		var arr []int

		for j := 0; j < n; j++ {
			var x int
			fmt.Scan(&x)
			arr = append(arr, x)

		}

		// find the largest element in arr
		var min, max int
		for i, e := range arr {
			if i == 0 || e < min {
				min = e
			}
			if i == 0 || e > max {
				max = e
			}
		}

		//fmt.Println(arr, min, max)
		fmt.Println(max - min)
	}
}
