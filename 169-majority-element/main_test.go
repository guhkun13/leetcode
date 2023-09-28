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
			desc:  "tc1",
			input: []int{3, 2, 3},
			want:  3,
		},
		{
			desc:  "tc2",
			input: []int{2, 2, 1, 1, 1, 2, 2},
			want:  2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.want, majorityElement(tC.input))
		})
	}
}
