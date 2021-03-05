package access

import (
	"fmt"
	"math"
)

var INF = math.Inf(-1)

type DataItem struct {
	Data int
}

func NewDataItem(data int) DataItem{
	return DataItem{
		data,
	}
}

type Node234 struct {
	Parent *Node234
	Children []Node234
	NumItems int
	DataItems []DataItem
	IsLeaf bool
}

func NewNode234() Node234{
	return Node234{
		nil,
		make([]Node234, 4),
		0,
		[]DataItem{{int(INF)},
			{int(INF)},
			{int(INF)}},
		true,
	}
}

type Tree234 struct {
	Root *Node234
}

func NewTree234() Tree234{
	return Tree234{
		Root: nil,
	}
}

func (t *Tree234) Insert(key int) error{
	currentNode := t.Root
	if t.Root == nil {
		rootNode := NewNode234()
		rootNode.InsertItem(key)
		t.Root = &rootNode
	} else {
			if currentNode.isFull() {
		currentNode.splitNode(key)
	}
	currentNode.InsertItem(key)
	}
	return nil
}

func (n *Node234) InsertItem(key int) error{
	n.NumItems++
	newKey := DataItem{key}
	for i := 2; i>= 0; i-- {
		if n.DataItems[i].Data == int(INF) {
			continue
		} else {
			currentKey := n.DataItems[i].Data
			fmt.Print(currentKey)

			if newKey.Data < currentKey {
				n.DataItems[i + 1].Data = n.DataItems[i].Data //Shift right
			} else {
				n.DataItems[i + 1].Data = key

				return nil
			}
		}
	}
	n.DataItems[0] = newKey
	return nil
}

func (n *Node234) isFull() bool {
	if n == nil {
		return false
	}
	if len(n.DataItems) >= 3 {
		return true
	}
	return false
}

func (n *Node234) splitNode(key int) error {
/*
	var parent, child2, child3 Node234
	var dataItem1, dataItem2 DataItem
	var itemIndex int
	newRight := NewNode234()
*/
	/*if this current node is a root
	 * make a new node, parent is root and connect to parent
	*/
	return nil
}

func (n *Node234) RemoveItem() *DataItem {
	//remove the largest item
	var retDataItem DataItem
	for i := len(n.DataItems)-1; i >= 0; i-- {
		retDataItem = n.DataItems[i]
		n.DataItems[i].Data = int(INF)
		n.NumItems--
		return &retDataItem
	}
	return nil
}