package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	result := 0

	dat, err := os.ReadFile("./input.txt")
	check(err)

	s := string(dat)

	target := "XMAS"
	target2 := "SAMX"

	r := strings.Split(s, "\n")
	for i, v := range r {
		for j := range v {
			rightD, leftD, bottom, right := "", "", "", ""

			if j+4 <= len(v) {
				right = v[j : j+4]
			}

			if i+3 < len(r) {
				bottom = string(r[i][j]) + string(r[i+1][j]) + string(r[i+2][j]) + string(r[i+3][j])
			}

			if i+3 < len(r) && j+3 < len(r[i+3]) {
				rightD = string(r[i][j]) + string(r[i+1][j+1]) + string(r[i+2][j+2]) + string(r[i+3][j+3])
			}

			if i+3 < len(r) && j-3 >= 0 {
				leftD = string(r[i][j]) + string(r[i+1][j-1]) + string(r[i+2][j-2]) + string(r[i+3][j-3])
			}

			if right == target || right == target2 {
				result++
			}
			if bottom == target || bottom == target2 {
				result++
			}
			if rightD == target || rightD == target2 {
				result++
			}
			if leftD == target || leftD == target2 {
				result++
			}
		}
	}
	fmt.Println(result)
}
