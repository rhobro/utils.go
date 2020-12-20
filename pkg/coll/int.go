package coll

func IndexInt(s []int, x int) int {
	for i, e := range s {
		if x == e {
			return i
		}
	}
	return -1
}

func ContainsInt(s []int, x int) bool {
	return IndexInt(s, x) > -1
}

func RemoveInt(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
