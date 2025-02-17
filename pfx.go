package domaintools

// HasHttpPrefix whether the value has a http prefix (regardless of the size of the characters - lowercase letters)
func HasHttpPrefix(str string) bool {
	return len(str) > 3 && (str[0] == 'h' || str[0] == 'H') &&
		(str[1] == 't' || str[1] == 'T') &&
		(str[2] == 't' || str[2] == 'T') &&
		(str[3] == 'p' || str[3] == 'P')
}

// HasHttpsPrefix whether the value has a https prefix (regardless of the size of the characters - lowercase letters)
func HasHttpsPrefix(str string) bool {
	return len(str) > 4 && (str[0] == 'h' || str[0] == 'H') &&
		(str[1] == 't' || str[1] == 'T') &&
		(str[2] == 't' || str[2] == 'T') &&
		(str[3] == 'p' || str[3] == 'P') &&
		(str[3] == 's' || str[3] == 's')
}

// HasIdnMarker optimized check if the str has a marker indicating that the domain is of the IDN type
// (regardless of the size of the characters - lowercase letters)
// however, this will not work for FQDNs that consist of several hostnames, and only one of them is of the IDN type,
// so in this situation it is better to use IsIdn
func HasIdnMarker(str string) bool {

	if len(str) < 4 {
		return false
	}
	step := 0
	for _, c := range str {
		switch step {
		case 0:
			if c == 'x' || c == 'X' {
				step = 1
			}
			break
		case 1:
			if c == 'n' || c == 'N' {
				step = 2
			} else {
				step = 0
			}
			break
		case 2:
			if c == '-' {
				step = 3
			} else {
				step = 0
			}
			break
		case 3:
			if c == '-' {
				return true
			}
			break

		}

	}

	return false
}

// CutWwwPrefix remove a www prefix from string (regardless of the size of the characters - lowercase letters)
func CutWwwPrefix(str string) string {
	if HasWwwPrefix(str) {
		return str[4:]
	}
	return str
}

// HasWwwPrefix whether the value has a www prefix  (regardless of the size of the characters - lowercase letters)
func HasWwwPrefix(str string) bool {
	// z.www.pl
	// 12345678
	// must be longer than 7 characters
	return len(str) > 7 &&
		(str[0] == 'w' || str[0] == 'W') &&
		(str[1] == 'w' || str[1] == 'W') &&
		(str[2] == 'w' || str[2] == 'W') &&
		str[3] == '.'
}
