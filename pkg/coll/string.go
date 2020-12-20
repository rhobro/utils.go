package coll

func IndexStr(s []string, x string) int {
	for i, e := range s {
		if x == e {
			return i
		}
	}
	return -1
}

func ContainsStr(s []string, x string) bool {
	return IndexStr(s, x) > -1
}

func RemoveStr(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
