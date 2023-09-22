package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		val  int
		want int
	}{
		{
			desc: "tc-1",
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: 2,
		},
		{
			desc: "tc-2",
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			want: 5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.want, RemoveElement(tC.nums, tC.val))
		})
	}
}
