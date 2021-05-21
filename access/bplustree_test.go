package access

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBPlusTree(t *testing.T) {

	tree := NewBPlusTree(4)
	tree.insert(5)
	tree.insert(3)
	tree.insert(8)
	assert.Equal(t, tree.Root.Keys, []int{3,5,8})
}


func TestBeyondFanout(t *testing.T) {
	tree := NewBPlusTree(3)
	tree.insert(5)
	tree.insert(3)
	res := tree.insert(8)
	fmt.Print(res)

}