package vlc

import (
	"strings"
	"unicode"
)

func Encode(str string) string {
	str = prepareText(str)

	// some text -> 10010101
	// split binary by chinks (8) 1010100010101010101010 -> 10101000 10101010 10101000 (bits -> bytes)
	// binary -> hex
	// return hexChinksStr
	return ""
}

// prepareText prepares text ot be fit for encode:
// changes upper case letters to : ! + lower case letter
func prepareText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}
