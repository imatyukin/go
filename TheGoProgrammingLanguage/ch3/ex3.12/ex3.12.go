// Упражнение 3.12.
// Напишите функцию, которая сообщает, являются ли две строки анаграммами одна другой, т.е. состоят ли они из одних и
// тех же букв в другом порядке.
package anagram

import (
	"strings"
	"sort"
)

func isAnagram(s1, s2 string) bool {
	if s1 == s2 || len(s1) != len(s2) {
		return false
	}

	s1 = SortString(s1)
	s2 = SortString(s2)
	if s1 == s2 {
		return true
	}
	return false
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
