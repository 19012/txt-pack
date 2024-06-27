package vlc

import (
	"reflect"
	"testing"
)

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

func TestNewHexChunks(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want HexChunks
	}{
		{
			name: "base test",
			str:  "20 30 3C 18",
			want: HexChunks{"20", "30", "3C", "18"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexChunks(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHexChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunk_toBin(t *testing.T) {
	tests := []struct {
		name string
		hc   HexChunk
		want BinChunk
	}{
		{
			name: "base test1",
			hc:   HexChunk("2F"),
			want: BinChunk("00101111"),
		},
		{
			name: "base test2",
			hc:   HexChunk("80"),
			want: BinChunk("10000000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hc.toBin(); got != tt.want {
				t.Errorf("HexChunk_toBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_ToBin(t *testing.T) {
	tests := []struct {
		name string
		hcs  HexChunks
		want BinChunks
	}{
		{
			name: "base test",
			hcs:  HexChunks{"2F", "80"},
			want: BinChunks{"00101111", "10000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hcs.ToBin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexChunks.ToBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
