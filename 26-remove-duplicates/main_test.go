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
			desc:  "t1",
			input: []int{1, 1, 2},
			want:  2,
		},
		{
			desc:  "t1",
			input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want:  5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.want, removeDuplicates(tC.input))
		})
	}
}
