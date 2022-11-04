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
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	// write a loop to read the test cases
	for i := 0; i < t; i++ {
		// read the number of elements
		var n int
		fmt.Fscan(reader, &n)

		// find the largest element in stream of input
		var min, max int64
		for j := 0; j < n; j++ {
			var x int64
			fmt.Fscan(reader, &x)

			if j == 0 || x < min {
				min = x
			}
			if j == 0 || x > max {
				max = x
			}
		}
		fmt.Fprintln(writer, max-min)
	}
}
