package access

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTree234(t *testing.T) {
	tree := NewTree234()
	tree.Insert(50)
	assert.Equal(t, tree.Root.DataItems[0].Data, 50)

}

func TestNewNode(t *testing.T) {
	tree := NewTree234()
	tree.Insert(50)
	tree.Insert(40)
}

func TestRemoveDataItem(t *testing.T){
	rootNode := Node234{
			nil,
			nil,
			3,
			[]DataItem{
				DataItem{40},
				DataItem{50},
				DataItem{60},
			},
			true,
	}
	tree := Tree234{
		&rootNode,
	}
	res := tree.Root.RemoveItem()
	assert.Equal(t, res.Data, 60)
	assert.Equal(t, tree.Root.NumItems, 2)
}

