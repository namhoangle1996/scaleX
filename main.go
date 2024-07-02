package main

import (
	"math"
)

func init() {
}

func main() {
}

func sol1(n int) (res int64) {
	var resFloat float64
	for i := 1; i <= n; i++ {
		resFloat += math.Pow(float64(i), 2)
	}
	return int64(resFloat)
}

/*
 Write a program that displays the result of total (S) below:
            S = 1^2 + 2^2 + 3^2 + … + N^2

            Test case 1: N = 10 Test case 2: N = 10^4, Test case 3: N = 10^6 , …
*
*/

/*
Write a program that displays the last 6 digits of S (Fibonacci) below:

	F1 = 0
	F2 = 1
	F3 = F1 + F2

S = F1 + F2 + F3 + … + Fn
Test case 1: N = 10 Test case 2: N = 10^4, Test case 3: N = 10^6 , …

Examples: if S = 123456789 then we display 456789

	if S = 1234 then we display 1234
*/
func sol2(n int) (res []int64) {
	var hmapF = make(map[int]int, n)
	hmapF[0] = 1
	hmapF[1] = 1

	var fSlice []int64
	fSlice = append(fSlice, int64(hmapF[0]), int64(hmapF[1]))

	for i := 2; i < n; i++ {
		hmapF[i] = hmapF[i-1] + hmapF[i-2]
		fSlice = append(fSlice, int64(hmapF[i]))
	}

	if len(fSlice) <= 6 {
		return fSlice
	}

	return fSlice[len(fSlice)-6:]
}
