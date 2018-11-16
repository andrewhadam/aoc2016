/*
Learning Go - AoC - Day1 - 2016
*/
package main

import "fmt"
import "strings"
import "regexp"
import "strconv"
import "math"

// AdjustLocation comment
func add(t int, x int, y int, inc int) (int, int) {
	if t == 0 {
		y += inc
	} else if t == 1 {
		x += inc
	} else if t == 2 {
		y -= inc
	} else {
		x -= inc
	}
	return x, y
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	var input = "R3, L5, R2, L1, L2, R5, L2, R2, L2, L2, L1, R2, L2, R4, R4, R1, L2, L3, R3, L1, R2, L2, L4, R4, R5, L3, R3, L3, L3, R4, R5, L3, R3, L5, L1, L2, R2, L1, R3, R1, L1, R187, L1, R2, R47, L5, L1, L2, R4, R3, L3, R3, R4, R1, R3, L1, L4, L1, R2, L1, R4, R5, L1, R77, L5, L4, R3, L2, R4, R5, R5, L2, L2, R2, R5, L2, R194, R5, L2, R4, L5, L4, L2, R5, L3, L2, L5, R5, R2, L3, R3, R1, L4, R2, L1, R5, L1, R5, L1, L1, R3, L1, R5, R2, R5, R5, L4, L5, L5, L5, R3, L2, L5, L4, R3, R1, R1, R4, L2, L4, R5, R5, R4, L2, L2, R5, R5, L5, L2, R4, R4, L4, R1, L3, R1, L1, L1, L1, L4, R5, R4, L4, L4, R5, R3, L2, L2, R3, R1, R4, L3, R1, L4, R3, L3, L2, R2, R2, R2, L1, L4, R3, R2, R2, L3, R2, L3, L2, R4, L2, R3, L4, R5, R4, R1, R5, R3"

	var whereami = 0
	var countx int
	var county int

	var slice = strings.Split(input, ",")

	m := make(map[string]int)
	var n []string

	r, _ := regexp.Compile("R(\\d+)")
	l, _ := regexp.Compile("L(\\d+)")

	for _, num := range slice {
		fmt.Println("Starting the loop")
		fmt.Println(countx)
		fmt.Println(county)
		fmt.Println(num)

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

			countx, county = add(whereami, countx, county, i)
			if Contains(n, strconv.Itoa(countx)+","+strconv.Itoa(county)) {
				fmt.Println("NEAT!A")
				break
			}

			n = append(n, strconv.Itoa(countx)+","+strconv.Itoa(county))

			fmt.Println("Post adjustment")
			fmt.Println(countx)
			fmt.Println(county)

			if _, ok := m[strconv.Itoa(countx)+","+strconv.Itoa(county)]; ok {
				fmt.Println("FOUND A MATCH A")
				//fmt.Println(val)
				fmt.Println(strconv.Itoa(countx) + "," + strconv.Itoa(county))
				break
			} else {
				m[strconv.Itoa(countx)+","+strconv.Itoa(county)] = 1
			}
		} else if l.MatchString(num) {

			b, err2 := strconv.Atoi(l.FindStringSubmatch(strings.Trim(num, " "))[1])
			if err2 != nil {
				fmt.Println(err2)
			}

			if whereami == 0 {
				whereami = 3
			} else {
				whereami--
			}

			countx, county = add(whereami, countx, county, b)

			if Contains(n, strconv.Itoa(countx)+","+strconv.Itoa(county)) {
				fmt.Println("NEAT!B")
				break
			}
			n = append(n, strconv.Itoa(countx)+","+strconv.Itoa(county))

			if _, ok := m[strconv.Itoa(countx)+","+strconv.Itoa(county)]; ok {
				fmt.Println("FOUND A MATCH B")
				//fmt.Println(val)
				fmt.Println(strconv.Itoa(countx) + "," + strconv.Itoa(county))
				break
			} else {
				m[strconv.Itoa(countx)+","+strconv.Itoa(county)] = 1
			}
		} else {
			fmt.Println("something really bad happened")

		}
		//		if num == " R5" {
		//		break
		//}
	}

	fmt.Println(math.Abs(float64(county)) + math.Abs(float64(countx)))
	fmt.Println(m)
}
