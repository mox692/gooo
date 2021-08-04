package unicode_substr

// UnicodeSubstr get string,
// and return first n words.(this word is Unicode 4byte code point).
// If n is over str's Unicode len, will return str[:].
func UnicodeSubstr(str string, n int) string {
	if uLen := len([]rune(str)); uLen > n {
		return string([]rune(str)[:n])
	}
	return str
}
