package golockserver

// FindArrElement finds the index of value's index and returns it or -1, otherwise
func FindArrElement(data []string, value string) int {
	for k, v := range data {
		if v == value {
			return k
		}
	}

	return -1
}
