package checkers

func StringItemExists(array []string, value string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}

	return false
}

func IntegerItemExists(array []int, value int) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}

	return false
}

func StringToBool(str string) bool {
	if str == "true" {
		return true
	}

	return false
}
