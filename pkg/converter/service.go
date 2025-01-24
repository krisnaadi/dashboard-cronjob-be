package converter

func BoolToInt(value bool) int8 {
	if value {
		return 1
	}

	return 0
}
