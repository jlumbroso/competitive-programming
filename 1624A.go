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

		// find the largest element in stream of input
		var min, max int
		for i := 0; i < n; i++ {
			var x int
			fmt.Scan(&x)

			if i == 0 || x < min {
				min = x
			}
			if i == 0 || x > max {
				max = x
			}
		}

		//fmt.Println(arr, min, max)
		fmt.Println(max - min)
	}
}
