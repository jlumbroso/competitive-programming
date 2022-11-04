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
		// read the length of the word
		var n int
		fmt.Fscan(reader, &n)

		var prevBigram string
		fmt.Fscan(reader, &prevBigram)
		fmt.Fprintln(writer, "<(", prevBigram, ")>")

		// find the largest element in stream of input
		for j := 0; j < n-2; j++ {
			var bigram string
			fmt.Fscan(reader, &bigram)

			// we found not consecutive bigrams so print full previous bigram
			if prevBigram[1] != bigram[0] {
				fmt.Fprintln(writer, prevBigram)
			} else {
				fmt.Fprintln(writer, string(prevBigram[0]))
			}
			prevBigram = bigram
		}
		fmt.Fprintln(writer, "//", t)
	}
}
