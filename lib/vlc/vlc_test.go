package vlc

import "testing"

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
