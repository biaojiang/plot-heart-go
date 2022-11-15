package main

import (
	"fmt"
	"math"
	"strings"
)

const PROFESSION string = "Engineer"

func main() {
	print_heart()
}

func print_heart() {

	for y := 15; y > -15; y-- {
		str_line := make([]string, 0) // print line-by-line
		y1 := float64(y)
		for x := -30; x < 30; x++ {
			x1 := float64(x)
			if math.Pow(math.Pow(x1*0.05, 2)+math.Pow(y1*0.1, 2)-1, 3)-math.Pow(x1*0.05, 2)*math.Pow(y1*0.1, 3) <= 0 {
				i := (x - y) % 8
				if i < 0 {
					i += 8
				}
				str_line = append(str_line, PROFESSION[i:i+1])
			} else {
				str_line = append(str_line, " ")
			}
		}
		fmt.Println(strings.Join(str_line, "")) // print one line
	}
}
