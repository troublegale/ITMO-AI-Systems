package main

import "slices"

func ModifyWithConstraint(s []int, c int, cond func(int, int) bool) []int {
	m := make([]int, 0)
	for _, n := range s {
		if cond(n, c) {
			m = append(m, n)
		}
	}
	return m
}

func GreaterThan(n int, c int) bool {
	return n > c
}

func LessThan(n int, c int) bool {
	return n < c
}

func Equals(n int, c int) bool {
	return n == c
}

func removeMatches(a, b []int) []int {
	var res []int
	for _, n := range a {
		if !slices.Contains(b, n) {
			res = append(res, n)
		}
	}
	return res
}
