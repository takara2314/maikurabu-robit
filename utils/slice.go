package utils

// StrSliceContains finds that target str is included in target str.
func StrSliceContains(s []string, target string) bool {
	for _, str := range s {
		if str == target {
			return true
		}
	}
	return false
}
