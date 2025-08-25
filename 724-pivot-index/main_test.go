package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "valid",
			input: []int{1, 7, 3, 6, 5, 6},
			want:  3,
		},
		{
			name:  "0",
			input: []int{2, 1, -1},
			want:  0,
		},
		{
			name:  "-1",
			input: []int{1, 2, 3},
			want:  -1,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			fmt.Println(tC.name)
			assert.Equal(t, tC.want, pivotIndex2(tC.input))
		})
	}

}
