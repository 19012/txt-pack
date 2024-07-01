package vlc

import (
	"strings"
	"unicode"
)

type encodingTable map[rune]string

func Encode(str string) string {
	str = prepareText(str)

	bStr := encodeBin(str)

	chunks := splitByChunks(bStr, chunkSize)

	return chunks.ToHex().ToString()
}

func Decode(str string) string {
	bStr := NewHexChunks(str).ToBin().ToString()

	dTree := getEncodingTable().DecodintTree()

	return returnUpperCaseLetterals(dTree.Decode(bStr))
}

// prepareText prepares text ot be fit for encode:
// changes <upper case letter> to : ! + <lower case letter>
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

// returnUpperCaseLetterals is opposite to prepareText, it prepares decoded test ot export:
// chnages <spase> + ! + <lower case letter> to <upper case letter>
func returnUpperCaseLetterals(str string) string {
	var buf strings.Builder

	var isCapital bool

	for _, ch := range str {
		if isCapital {
			if !unicode.IsLetter(ch) {
				buf.WriteRune(ch)
			} else {
				buf.WriteRune(unicode.ToUpper(ch))
			}
			isCapital = false

			continue
		}
		if ch == '!' {
			isCapital = true
			continue
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

// encodeBin encodes runes into binary codes string without spaces
func encodeBin(str string) string {
	var buf strings.Builder

	for _, r := range str {
		buf.WriteString(runeToBin(r))
	}

	return buf.String()
}

func runeToBin(ch rune) string {
	table := getEncodingTable()

	res, ok := table[ch]
	if !ok {
		panic("unknown character: " + string(ch))
	}

	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ':  "11",
		't':  "1001",
		'n':  "10000",
		's':  "0101",
		'r':  "01000",
		'd':  "00101",
		'!':  "001000",
		'c':  "000101",
		'm':  "000011",
		'g':  "0000100",
		'b':  "0000010",
		'v':  "00000001",
		'k':  "0000000001",
		'q':  "000000000001",
		'e':  "101",
		'o':  "10001",
		'a':  "011",
		'i':  "01001",
		'h':  "0011",
		'l':  "001001",
		'u':  "00011",
		'f':  "000100",
		'\n': "00000111",
		'p':  "0000101",
		'w':  "0000011",
		'y':  "0000001",
		'j':  "000000001",
		'x':  "00000000001",
		'z':  "000000000000",
	}
}
