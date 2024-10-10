package utils

import "strings"

func GlogToRegex(glog string) string {
	return glogToRegex(glog, true, true, "#")
}

func glogToRegex(glob string, strictLeadingDot bool, strictWildcardSlash bool, delimiter string) string {
	firstByte := true
	escaping := false
	inCurlies := 0
	var regex strings.Builder
	for i := 0; i < len(glob); i++ {
		car := glob[i]

		if firstByte && strictLeadingDot && car != '.' {
			regex.WriteString("(?=[^\\.])")
		}
		firstByte = car == '/'

		if firstByte && strictWildcardSlash && i+2 < len(glob) && glob[i+1] == '*' && glob[i+2] == '*' && (i+3 >= len(glob) || glob[i+3] == '/') {
			car = '/'
			regex.WriteString("(?:[^/]+/")
			if i+3 >= len(glob) {
				regex.WriteString("?)")
			}
			regex.WriteString(")*")
			i += 2 + btoi(i+3 < len(glob) && glob[i+3] == '/')
			continue
		}

		switch car {
		case '*':
			if escaping {
				regex.WriteString("\\*")
			} else {
				if strictWildcardSlash {
					regex.WriteString("[^/]*")
				} else {
					regex.WriteString(".*")
				}
			}
		case '?':
			if escaping {
				regex.WriteString("\\?")
			} else {
				if strictWildcardSlash {
					regex.WriteString("[^/]")
				} else {
					regex.WriteString(".")
				}
			}
		case '{':
			if escaping {
				regex.WriteString("\\{")
			} else {
				inCurlies++
				regex.WriteString("(")
			}
		case '}':
			if inCurlies > 0 && !escaping {
				inCurlies--
				regex.WriteString(")")
			} else {
				regex.WriteString("}")
			}
		case ',':
			if inCurlies > 0 && !escaping {
				regex.WriteString("|")
			} else {
				regex.WriteString(",")
			}
		case '\\':
			if escaping {
				regex.WriteString("\\\\")
				escaping = false
			} else {
				escaping = true
				continue
			}
		case '.', '(', ')', '|', '+', '^', '$':
			regex.WriteString("\\" + string(car))
		default:
			regex.WriteString(string(car))
		}
		escaping = false
	}

	return delimiter + "^" + regex.String() + "$" + delimiter
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
