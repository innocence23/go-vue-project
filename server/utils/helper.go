package utils

import "strconv"

func IntSliceToStringSlice(data []int) []string {
	res := make([]string, len(data))
	for i, x := range data {
		res[i] = strconv.Itoa(x)
	}
	return res
}
