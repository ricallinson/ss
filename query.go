package main

import (
	"os"
	"strings"
)

// Currently support; "a b c OR d NOT w x y OR z"
func Query(q string, o *os.File, c int) func(string) int {
	matches := strings.Split(q, " ")
	ands := []string{}
	nots := []string{}
	andOrs := [][]string{}
	notOrs := [][]string{}
	isAnd := true
	isOr := false
	andOrsIndex := 0
	notOrsIndex := 0

	for i, v := range matches {
		if v == "OR" {
			isOr = true
		} else if v == "NOT" {
			isAnd = false
		} else {
			if i+1 < len(matches) && matches[i+1] == "OR" {
				isOr = true
			}
			if isOr && i-1 > 0 && matches[i-1] != "OR" {
				if isAnd {
					andOrsIndex = len(andOrs)
				} else {
					notOrsIndex = len(notOrs)
				}
			}
			if isAnd {
				if isOr {
					if len(andOrs) == andOrsIndex {
						andOrs = append(andOrs, []string{})
					}
					andOrs[andOrsIndex] = append(andOrs[andOrsIndex], v)
					isOr = false
				} else {
					ands = append(ands, v)
				}
			} else {
				if isOr {
					if len(notOrs) == notOrsIndex {
						notOrs = append(notOrs, []string{})
					}
					notOrs[notOrsIndex] = append(notOrs[notOrsIndex], v)
					isOr = false
				} else {
					nots = append(nots, v)
				}
			}
		}
	}

	// fmt.Println("ANDs", ands, "AND ORs", andOrs, "NOTs", nots, "NOT ORs", notOrs)

	andsTotal := len(ands) + len(andOrs)
	notsTotal := len(nots) + len(notOrs)

	return func(line string) int {
		if andsTotal+notsTotal == 0 {
			return 0
		}
		andsCheck := 0
		notsCheck := 0
		for _, and := range ands {
			if strings.Contains(line, and) {
				andsCheck++
			}
		}
		for _, not := range nots {
			if strings.Contains(line, not) {
				notsCheck++
			}
		}
		for _, ors := range andOrs {
			for _, or := range ors {
				if strings.Contains(line, or) {
					andsCheck++
					break
				}
			}
		}
		for _, ors := range notOrs {
			for _, or := range ors {
				if strings.Contains(line, or) {
					notsCheck++
					break
				}
			}
		}

		if andsCheck == andsTotal && notsCheck == 0 {
			if o != nil && (c == -1 || c > 0) {
				o.WriteString(line + "\n")
				if c > 0 {
					c--
				}
				if c == 0 {
					os.Exit(0)
				}
			}
			return 1
		}
		return 0
	}
}
