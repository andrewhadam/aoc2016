/*
Learning Go - AoC - Day1 - 2016
*/
package main

import "fmt"
import "strings"
import "regexp"
import "strconv"

func main() {
	var input = "R3, L5, R2, L1, L2, R5, L2, R2, L2, L2, L1, R2, L2, R4, R4, R1, L2, L3, R3, L1, R2, L2, L4, R4, R5, L3, R3, L3, L3, R4, R5, L3, R3, L5, L1, L2, R2, L1, R3, R1, L1, R187, L1, R2, R47, L5, L1, L2, R4, R3, L3, R3, R4, R1, R3, L1, L4, L1, R2, L1, R4, R5, L1, R77, L5, L4, R3, L2, R4, R5, R5, L2, L2, R2, R5, L2, R194, R5, L2, R4, L5, L4, L2, R5, L3, L2, L5, R5, R2, L3, R3, R1, L4, R2, L1, R5, L1, R5, L1, L1, R3, L1, R5, R2, R5, R5, L4, L5, L5, L5, R3, L2, L5, L4, R3, R1, R1, R4, L2, L4, R5, R5, R4, L2, L2, R5, R5, L5, L2, R4, R4, L4, R1, L3, R1, L1, L1, L1, L4, R5, R4, L4, L4, R5, R3, L2, L2, R3, R1, R4, L3, R1, L4, R3, L3, L2, R2, R2, R2, L1, L4, R3, R2, R2, L3, R2, L3, L2, R4, L2, R3, L4, R5, R4, R1, R5, R3"

	var whereami = 0
	var countx = 0
	var county = 0

	var slice = strings.Split(input, ",")

	r, _ := regexp.Compile("R(\\d+)")
	l, _ := regexp.Compile("L(\\d+)")

	for _, num := range slice {

		if r.MatchString(num) {
			i, err := strconv.Atoi(r.FindStringSubmatch(strings.Trim(num, " "))[1])
			if err != nil {
				fmt.Println(err)
				return
			}

			if whereami == 3 {
				whereami = 0
			} else {
				whereami++
			}
			if whereami == 0 {
				county += i
			} else if whereami == 1 {
				countx += i
			} else if whereami == 2 {
				county -= i
			} else {
				countx -= i
			}

		} else if l.MatchString(num) {

			b, err2 := strconv.Atoi(l.FindStringSubmatch(strings.Trim(num, " "))[1])
			if whereami == 0 {
				whereami = 3
			} else {
				whereami--
			}
			if whereami == 0 {
				county += b
			} else if whereami == 1 {
				countx += b
			} else if whereami == 2 {
				county -= b
			} else {
				countx -= b
			}

			if err2 != nil {
				fmt.Println(err2)
			}
		} else {
			fmt.Println("something really bad happened")

		}
	}
	fmt.Println(strconv.Itoa(countx) + " R moves")
	fmt.Println(strconv.Itoa(county) + " L moves")
}
