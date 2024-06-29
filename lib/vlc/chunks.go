package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const chunkSize = 8

const sep = ' '

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

// ToString joins chunks into one line and returns as string
func (hcs HexChunks) ToString() string {
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

// ToString joins chunks into one line and returns as string
func (bcs BinChunks) ToString() string {
	switch len(bcs) {
	case 0:
		return ""
	case 1:
		return string(bcs[0])
	}
	var buf strings.Builder

	for _, chunk := range bcs {
		buf.WriteString(string(chunk))
	}

	return buf.String()
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

func NewHexChunks(str string) HexChunks {
	parts := strings.Split(str, string(sep))

	res := make(HexChunks, 0, len(parts))

	for _, p := range parts {
		res = append(res, HexChunk(p))
	}

	return res
}

func (hcs HexChunks) ToBin() BinChunks {
	res := make(BinChunks, 0, len(hcs))
	for _, chunk := range hcs {
		res = append(res, chunk.toBin())
	}
	return res
}

func (hc HexChunk) toBin() BinChunk {
	if num, err := strconv.ParseUint(string(hc), 16, chunkSize); err != nil {
		panic("can't parse hex chunks: " + err.Error())
	} else {
		return BinChunk(fmt.Sprintf("%08b", num))
	}
}
