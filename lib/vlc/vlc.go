package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const chunkSize = 8

type encodingTable map[rune]string

type BinChunk string

type BinChunks []BinChunk

type HexChunk string

type HexChunks []HexChunk

func (bcs BinChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bcs))

	for _, chunk := range bcs {
		res = append(res, chunk.toHex())
	}

	return res
}

func (bc BinChunk) toHex() HexChunk {
	if num, err := strconv.ParseUint(string(bc), 2, chunkSize); err != nil {
		panic("can't parse binary chink: " + err.Error())
	} else {
		res := strings.ToUpper(fmt.Sprintf("%x", num))

		if len(res) == 1 {
			res = "0" + res
		}

		return HexChunk(res)
	}
}

func (hcs HexChunks) ToString() string {
	const sep = ' '

	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}

	var buf strings.Builder

	buf.WriteString(string(hcs[0]))

	for _, chunk := range hcs[1:] {
		buf.WriteRune(sep)
		buf.WriteString(string(chunk))
	}

	return buf.String()
}

func Encode(str string) string {
	str = prepareText(str)

	bStr := encodeBin(str)

	chunks := splitByChunks(bStr, chunkSize)

	return chunks.ToHex().ToString()
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

// encodeBin encodes runes into binary codes string without spaces
func encodeBin(str string) string {
	var buf strings.Builder

	for _, r := range str {
		buf.WriteString(runeToBin(r))
	}

	return buf.String()
}

func runeToBin(r rune) string {
	if res, ok := getEncodingTable()[r]; !ok {
		panic("unknown character: " + string(r))
	} else {
		return res
	}
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

// splitByChunks splits binary string by chunks with given chunk size:
func splitByChunks(bStr string, chunkSize int) BinChunks {
	strLen := utf8.RuneCountInString(bStr)

	chunksCount := strLen / chunkSize

	if strLen%chunkSize != 0 {
		chunksCount++
	}

	var buf strings.Builder

	res := make(BinChunks, 0, chunksCount)

	for i, r := range bStr {
		buf.WriteString(string(r))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()

		lastChunk += strings.Repeat("0", chunkSize-buf.Len())

		res = append(res, BinChunk(lastChunk))
	}

	return res
}
