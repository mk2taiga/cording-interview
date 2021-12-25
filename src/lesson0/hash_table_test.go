package lesson0

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashTable(t *testing.T) {
	tests := [...]struct {
		cap       int
		addKey    []int
		addValue  []int
		removeKey []int
		searchKey []int
		want      []int
		wantDump  []string
	}{
		{
			cap:       5,
			addKey:    []int{1, 2, 3, 4, 5, 6},
			addValue:  []int{10, 20, 30, 40, 50, 60},
			removeKey: []int{1, 3},
			searchKey: []int{1, 2, 10, 3, 4, 5, 6},
			want:      []int{-1, 20, -1, -1, 40, 50, 60},
			wantDump: []string{
				"key=5,val=50",
				"key=6,val=60",
				"key=2,val=20",
				"key=4,val=40",
			},
		},
	}

	for _, tt := range tests {
		hash := NewHashTable(tt.cap)
		for i, key := range tt.addKey {
			assert.True(t, hash.Add(key, tt.addValue[i]))
		}
		for _, key := range tt.removeKey {
			assert.True(t, hash.Remove(key))
		}
		for i, key := range tt.searchKey {
			assert.Equal(t, tt.want[i], hash.Search(key))
		}
		assert.Equal(t, tt.wantDump, hash.Dump())
	}
}
