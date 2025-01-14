package utils

import "sort"

func InStringSlice(key string, s []string) bool {
	for _, v := range s {
		if key == v {
			return true
		}
	}
	return false
}

// CompareIntSlice 比较两个int切片是否相等
func CompareIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Ints(a)
	sort.Ints(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// SliceIntersect 两切片交集
func SliceIntersect(a []string, b []string) []string {
	inter := make([]string, 0)
	mp := make(map[string]bool)

	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			inter = append(inter, s)
		}
	}

	return inter
}
