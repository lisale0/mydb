package access

import (
	"fmt"
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


func TestDisconnectChild(t *testing.T) {
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
	child1 := tree.Root.DisconnectChild(1)
	assert.Equal(t, child1.DataItems[0].Data, 50)
	assert.Equal(t, tree.Root.DataItems[1].Data, int(INF))
	child2 := tree.Root.DisconnectChild(2)
	assert.Equal(t, child2.DataItems[0].Data, 60)
	assert.Equal(t, tree.Root.DataItems[2].Data, int(INF))
}

func TestConnectChild(t *testing.T) {
	rootNode := Node234{
		nil,
		make([]Node234, 3),
		3,
		make([]DataItem, 3),
		true,
	}

	newNode := Node234{
		nil,
		make([]Node234, 3),
		3,
		[]DataItem{
			{567},
		},
		true,
	}
	rootNode.ConnectChild(0, &newNode)
	retParent, _ := newNode.GetParent()
	assert.Equal(t, rootNode.Children[0].DataItems, newNode.DataItems)
	assert.Equal(t, *retParent, rootNode)

}

func TestSplit(t *testing.T){
	tree := NewTree234()
	tree.Insert(50)
	tree.Insert(40)
	tree.Insert(60)
	fmt.Print(tree)
}