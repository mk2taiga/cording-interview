package lesson0

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayList(t *testing.T) {
	tests := [...]struct {
		addValue []int
		wantCap  int
		wantDump []int
	}{
		{
			addValue: []int{10, 20, 30, 40, 50, 60},
			wantCap:  8,
			wantDump: []int{10, 20, 30, 40, 50, 60, 0, 0},
		},
	}

	for _, tt := range tests {
		list := &ArrayList{}
		for _, v := range tt.addValue {
			assert.True(t, list.Add(v))
		}

		assert.Equal(t, tt.wantCap, list.Cap())
		assert.Equal(t, tt.wantDump, list.Dump())
	}
}
