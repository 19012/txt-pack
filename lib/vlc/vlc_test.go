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
		want BinChanks
	}{
		{
			name: "base test",
			args: args{
				bStr:      "001000100110100101",
				chunkSize: 8,
			},
			want: BinChanks{"00100010", "01101001", "01000000"},
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
