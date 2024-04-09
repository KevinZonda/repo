package utils

func NotEmptyOne(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	for _, v := range s {
		if v != "" {
			return v
		}
	}
	return ""
}
