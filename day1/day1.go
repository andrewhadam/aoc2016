package main

import "fmt"
import "strings"
import "regexp"
import "strconv"

/*
You've got N, E, S, W (0, 1, 2, 3)
We basically want to count the slope
0 -> + y
1 -> + x
2 -> - y
3 -> - x

Add x and why
*/
func main() {
	var input = "R3, L5, R2, L1, L2, R5, L2, R2, L2, L2, L1, R2, L2, R4, R4, R1, L2, L3, R3, L1, R2, L2, L4, R4, R5, L3, R3, L3, L3, R4, R5, L3, R3, L5, L1, L2, R2, L1, R3, R1, L1, R187, L1, R2, R47, L5, L1, L2, R4, R3, L3, R3, R4, R1, R3, L1, L4, L1, R2, L1, R4, R5, L1, R77, L5, L4, R3, L2, R4, R5, R5, L2, L2, R2, R5, L2, R194, R5, L2, R4, L5, L4, L2, R5, L3, L2, L5, R5, R2, L3, R3, R1, L4, R2, L1, R5, L1, R5, L1, L1, R3, L1, R5, R2, R5, R5, L4, L5, L5, L5, R3, L2, L5, L4, R3, R1, R1, R4, L2, L4, R5, R5, R4, L2, L2, R5, R5, L5, L2, R4, R4, L4, R1, L3, R1, L1, L1, L1, L4, R5, R4, L4, L4, R5, R3, L2, L2, R3, R1, R4, L3, R1, L4, R3, L3, L2, R2, R2, R2, L1, L4, R3, R2, R2, L3, R2, L3, L2, R4, L2, R3, L4, R5, R4, R1, R5, R3"

	var whereami = 0
	var countx = 0
	var county = 0

	var slice = strings.Split(input, ",")
	fmt.Println(slice)
	fmt.Println()

	for _, num := range slice {
		fmt.Println(num)
		//      ok, err := regexp.MatchString("R\\d+", num)
		r, _ := regexp.Compile("R(\\d+)")
		l, _ := regexp.Compile("L(\\d+)")

		if r.MatchString(num) {
			i, err := strconv.Atoi(r.FindStringSubmatch(strings.Trim(num, " "))[1])
			if whereami == 3 {
				whereami = 0
			} else {
				whereami += 1
			}
			fmt.Println("Right shark")
			fmt.Println(i)
			if whereami == 0 {
				countx += i
			} else if whereami == 1 {
				county += i
			} else if whereami == 2 {
				countx -= i
			} else {
				county -= i
			}
			if err != nil {
				fmt.Println(err)
			}
		} else if l.MatchString(num) {
			fmt.Println("break?")
			fmt.Println(num + "----my string")
			fmt.Println(strings.Trim(num, " "))
			b, err2 := strconv.Atoi(l.FindStringSubmatch(strings.Trim(num, " "))[1])
			if whereami == 0 {
				whereami = 3
			} else {
				whereami -= 1
			}
			fmt.Println("Left shark")

			if whereami == 0 {
				countx += b
			} else if whereami == 1 {
				county += b
			} else if whereami == 2 {
				countx -= b
			} else {
				county -= b
			}
			county += 1
			if err2 != nil {
				fmt.Println(err2)
			}
		} else {
			fmt.Println("something really bad happened")

		}
	}
	fmt.Println(strconv.Itoa(countx) + " R moves")
	fmt.Println(strconv.Itoa(county) + " L moves")
	fmt.Println(whereami)
}
