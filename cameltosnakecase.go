package piscine

func CamelToSnakeCase(s string) string {
	result := ""
	sentence := []rune(s)
	if IsUpper(sentence[len(s)-1]) {
		return s
	}
	for ind, elem := range sentence {
		if IsUpper(elem) && ind != 0 {
			if !IsUpper(sentence[ind+1]) && ind < len(sentence)-1 {
				result += "_"
			} else {
				return s
			}
		}
		result += string(elem)
	}
	return result
}

func IsUpper(l rune) bool {
	if l >= 'A' && l <= 'Z' {
		return true
	}
	return false
}
