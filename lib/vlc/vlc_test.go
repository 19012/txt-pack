package vlc

import (
	"reflect"
	"testing"
)

func Test_prepareText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "My name is Ted!",
			want: "!my name is !ted!",
		},
		{
			name: "empry input test",
			str:  "",
			want: "",
		},
		{
			name: "upper case only test",
			str:  "HELLO",
			want: "!h!e!l!l!o",
		},
		{
			name: "lower case only test",
			str:  "dream",
			want: "dream",
		},
		{
			name: "punctuation only test",
			str:  ".,!@#^yi",
			want: ".,!@#^yi",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = #{got}, want #{tt.want}")
			}
		})
	}
}

func Test_encodeBin(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "!ted",
			want: "001000100110100101",
		},
		{
			name: "empty input test",
			str:  "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBin(tt.str); got != tt.want {
				t.Errorf("encodeBin() = #{got}, want #{tt.want}")
			}
		})
	}
}

func Test_splitByChunks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}

	tests := []struct {
		name string
		args args
		want BinChunks
	}{
		{
			name: "base test",
			args: args{
				bStr:      "001000100110100101",
				chunkSize: 8,
			},
			want: BinChunks{"00100010", "01101001", "01000000"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinChunks
		want HexChunks
	}{
		{
			name: "base test",
			bcs:  BinChunks{"0101111", "10000000"},
			want: HexChunks{"2F", "80"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinChunks.ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Encode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); got != tt.want {
				t.Errorf("Encode() = #{got}, want #{tt.want}")
			}
		})
	}
}
