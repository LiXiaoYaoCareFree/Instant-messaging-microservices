package utils

func InList(list []string, key string) (ok bool) {
	for _, s := range list {
		if s == key {
			return true
		}
	}
	return false
}
