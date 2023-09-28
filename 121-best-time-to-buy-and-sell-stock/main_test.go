package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc  string
		input []int
		want  int
	}{
		{
			desc:  "tc-1",
			input: []int{7, 1, 5, 3, 6, 4},
			want:  5,
		},
		{
			desc:  "tc-2",
			input: []int{7, 6, 4, 3, 1},
			want:  0,
		},
		{
			desc:  "tc-3",
			input: []int{1, 2},
			want:  1,
		},
		{
			desc:  "tc-3",
			input: []int{2, 4, 1},
			want:  2,
		},
		{
			desc:  "tc-3",
			input: []int{2, 1, 2, 1, 0, 1, 2},
			want:  2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.want, MaxProfit(tC.input))
		})
	}
}
