package trails

func UintContains(s []uint, str uint) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
