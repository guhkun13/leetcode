package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  bool
	}{
		{
			desc:  "T1",
			input: "()",
			want:  true,
		},
		{
			desc:  "T2",
			input: "()[]{}",
			want:  true,
		},
		{
			desc:  "F1",
			input: "(]",
			want:  false,
		},
		{
			desc:  "F2",
			input: "",
			want:  false,
		},
		{
			desc:  "F3",
			input: "()]",
			want:  false,
		},
		{
			desc:  "F4",
			input: "(}",
			want:  false,
		},
		{
			desc:  "Adv-T1",
			input: "{[]}",
			want:  true,
		},
		{
			desc:  "Adv-F1",
			input: "((",
			want:  false,
		},
		{
			desc:  "Adv-F2",
			input: "){",
			want:  false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.want, IsValid(tC.input))
		})
	}
}
