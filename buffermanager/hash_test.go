package buffermanager

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBufHashTable_Hash(t *testing.T) {
	hashTable := NewBufHashTable()
	actualIdx := hashTable.Hash(118)
	idx := 118 % HTSIZE
	assert.Equal(t, idx, actualIdx)
}