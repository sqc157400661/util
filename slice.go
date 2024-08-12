package util

import (
	"sort"
	"strings"
)

func IsStringSliceEqual(slice1, slice2 []string) bool {
	// 比较长度
	if len(slice1) != len(slice2) {
		return false
	}
	sort.Strings(slice1)
	sort.Strings(slice2)
	// 比较对应位置的元素是否相同
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		} else {
			continue
		}
	}
	return true
}

func InStringSlice(slice []string, value string) bool {
	for _, s := range slice {
		if s == value {
			return true
		}
	}
	return false
}

func InIntSlice(slice []int, value int) bool {
	for _, s := range slice {
		if s == value {
			return true
		}
	}
	return false
}

func ToUpperStringSlice(slice []string) []string {
	var goal []string
	for _, s := range slice {
		goal = append(goal, strings.ToUpper(s))
	}
	return goal
}
