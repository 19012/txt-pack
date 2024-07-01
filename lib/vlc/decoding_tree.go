package vlc

import "strings"

type DecodingTree struct {
	Value rune
	Zero  *DecodingTree
	One   *DecodingTree
}

func (et encodingTable) DecodintTree() DecodingTree {
	res := DecodingTree{}

	for r, code := range et {
		res.Add(r, code)
	}

	return res
}

func (dt *DecodingTree) Add(r rune, code string) {
	curr := dt
	for _, b := range code {
		switch b {
		case '0':
			if curr.Zero == nil {
				curr.Zero = &DecodingTree{}
			}
			curr = curr.Zero
		case '1':
			if curr.One == nil {
				curr.One = &DecodingTree{}
			}
			curr = curr.One
		}
	}

	curr.Value = r
}

func (dt *DecodingTree) Decode(str string) string {
	var buf strings.Builder

	curr := dt

	for _, b := range str {
		if curr.Value != 0 {
			buf.WriteRune(curr.Value)
			curr = dt
		}

		switch b {
		case '0':
			curr = curr.Zero
		case '1':
			curr = curr.One
		}
	}

	if curr.Value != 0 {
		buf.WriteRune(curr.Value)
		curr = dt
	}

	return buf.String()
}
