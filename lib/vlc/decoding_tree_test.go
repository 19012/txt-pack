package vlc

import (
	"reflect"
	"testing"
)

func Test_encodinTable_DecodingTree(t *testing.T) {
	tests := []struct {
		name string
		et   encodingTable
		want DecodingTree
	}{
		{
			name: "base test",
			et: encodingTable{
				'a': "11",
				'b': "1001",
				'c': "0101",
			},
			want: DecodingTree{
				Zero: &DecodingTree{
					One: &DecodingTree{
						Zero: &DecodingTree{
							One: &DecodingTree{
								Value: 'c',
							},
						},
					},
				},
				One: &DecodingTree{
					Zero: &DecodingTree{
						Zero: &DecodingTree{
							One: &DecodingTree{
								Value: 'b',
							},
						},
					},
					One: &DecodingTree{
						Value: 'a',
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.DecodintTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodintTree() = #{got}, want #{tt.want}")
			}
		})
	}
}
