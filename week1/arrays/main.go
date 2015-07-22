package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	min := x[0]

	for _, value := range x {
		if value < min {
			min = value
		}
	}
	fmt.Println(min)

	elements := map[string]map[string]string{
		"CA": map[string]string{
			"name":    "California",
			"capital": "Sacramento",
		},
		"HI": map[string]string{
			"name":    "Hawaii",
			"capital": "Honolulu",
		},
	}
	if el, ok := elements["CA"]; ok {
		fmt.Println(el["name"], el["capital"])
	}
}

// elements["LA"] = "Louisiana"
// elements["VT"] = "Vermont"
// elements["OR"] = "Oregon"
// elements["CO"] = "Colorado"
// elements["NE"] = "NEVADA"
