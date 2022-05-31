package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{1, 5, -1, 6, 3, 2}, 3))
}

func maxSubArray(array []int, window int) int {
	count := 1
	j := 0
	max := 0
	sum := 0
	for i := 0; i < len(array); i++ {
		if count <= window {
			sum += array[i]
			count++
		} else {
			sum -= array[j]
			sum += array[i]
			j++
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
