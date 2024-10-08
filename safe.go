package domaintools

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

const acePrefix = "xn--"

func ToASCII(s string) (string, error) {
	if ascii(s) {
		return s, nil
	}
	labels := strings.Split(s, ".")
	for i, label := range labels {
		if !ascii(label) {
			a, err := encode(acePrefix, label)
			if err != nil {
				return "", err
			}
			labels[i] = a
		}
	}
	return strings.Join(labels, "."), nil
}

func ToUnicode(s string) (string, error) {
	if !strings.Contains(s, acePrefix) {
		return s, nil
	}
	labels := strings.Split(s, ".")
	for i, label := range labels {
		if strings.HasPrefix(label, acePrefix) {
			u, err := decode(label[len(acePrefix):])
			if err != nil {
				return "", err
			}
			labels[i] = u
		}
	}
	return strings.Join(labels, "."), nil
}

func ascii(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			return false
		}
	}
	return true
}

func SafeToIdn(s string) string {
	if !strings.Contains(s, acePrefix) {
		return s
	}

	str, err := ToUnicode(s)
	if err != nil {
		return s
	} else {
		return str
	}

}

func SafeToAscii(s string) string {
	if strings.Contains(s, acePrefix) {
		return s
	}

	str, err := ToASCII(s)
	if err != nil {
		return s
	} else {
		return str
	}
}

func IsIdn(s string) bool {
	return strings.HasPrefix(s, acePrefix)
}

func SafeAsciiDomainName(s string) string {
	return strings.ToLower(SafeToAscii(strings.Trim(s, " ")))

}

func GetDomainPartFromEmailAddress(s string) string {
	if pos := strings.Index(s, "@"); pos > 1 && pos+1 > len(s) {
		return s[pos+1:]
	} else {
		return ""
	}
}

func IsValidEmailAddress(emailAddress string) bool {
	return emailRegExp.MatchString(emailAddress)
}

func IsValidEmailAddressLowerCase(emailAddress string) bool {
	return emailRegExp.MatchString(emailAddress)
}

const emailPattern = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

var emailRegExp = regexp.MustCompile(emailPattern)