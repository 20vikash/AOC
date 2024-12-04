package main

import "fmt"

func part2(r []string) {
	cross1 := ""
	cross2 := ""
	result := 0

	for i, v := range r {
		for j := range v {
			cross1, cross2 = "", ""

			if v[j] == 'S' || v[j] == 'M' {
				if j+2 < len(v) && i+2 < len(r) {
					cross1 = string(v[j]) + string(r[i+1][j+1]) + string(r[i+2][j+2])
					cross2 = string(v[j+2]) + string(r[i+1][j+1]) + string(r[i+2][j])
				} else {
					continue
				}

				if (cross1 == "MAS" || cross1 == "SAM") && (cross2 == "MAS" || cross2 == "SAM") {
					result++
				}
			}
		}
	}

	fmt.Println(result)
}
