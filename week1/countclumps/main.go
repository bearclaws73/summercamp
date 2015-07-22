package main

import "fmt"

func countClumps(xs []int) int {

	var clumps int
	inClump := false
	for i := 1; i < len(xs); i++ {
		if inClump {
			if xs[i] != xs[i-1] {
				inClump = false
			}
		} else {
			if xs[i] == xs[i-1] {
				clumps++
				inClump = true
			}
		}
	}
	return clumps
}

func main() {

	clumps := countClumps([]int{1, 1, 3, 3, 1, 1})
	fmt.Println(clumps)

}
