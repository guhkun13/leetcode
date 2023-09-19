package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "valid",
			input: []string{"flower", "flow", "flight"},
			want:  "fl",
		},
		{
			name:  "empty",
			input: []string{"dog", "racecar", "car"},
			want:  "",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			fmt.Println(tC.name)
			assert.Equal(t, tC.want, MainFunc(tC.input))
		})
	}

}
