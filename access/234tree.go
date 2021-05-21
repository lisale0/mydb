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
	} else if currentNode.isFull() {
		t.Split(*currentNode)
	} else {
		currentNode.InsertItem(key)
	}
	return nil
}

func (t *Tree234) Split(n Node234) error {
	var parent *Node234
	var newRight Node234
	child2 := n.DisconnectChild(1)
	child3 := n.DisconnectChild(2)
	itemC := n.RemoveItem()
	itemB := n.RemoveItem()


	// if the current node is root, make a new root
	if n.IsRoot() {
		root := Node234{}
		parent = t.Root
		root.ConnectChild(0, &n)
	} else {
		parent, _ = n.GetParent() //n is not parent, get the parent
	}

	// deal with the parent
	itemIndex := parent.InsertItem(itemB.Data)
	numItems, _ := parent.GetNumItems()

	// move the parent's connection one child at a time
	for i:= numItems - 1; i > 0; i--{
		temp := parent.DisconnectChild(i)
		parent.ConnectChild(i+1, &temp)
	}
	parent.ConnectChild(0, &newRight)
	//deal with the newRight

	fmt.Print(itemC)
	fmt.Print(itemB)
	fmt.Print(child2)
	fmt.Print(child3)
	fmt.Print(itemIndex)
	fmt.Print(numItems)

	fmt.Print(parent)
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
	fill := 0
	for i := 0; i < len(n.DataItems); i++{
		if n.DataItems[i].Data > int(INF) {
			fill += 1
		}
	}
	if fill >= 2 {
		return true
	}

	return false
}



func (n *Node234) RemoveItem() *DataItem {
	//remove the largest item
	var retDataItem DataItem
	for i := len(n.DataItems)-1; i >= 0; i-- {
		retDataItem = n.DataItems[i]
		if retDataItem.Data == int(INF) {
			continue
		}

		n.DataItems[i].Data = int(INF)
		n.NumItems--
		return &retDataItem
	}
	return nil
}

//disconnect from node and return a new node that is to be a child at specified index
func (n *Node234) DisconnectChild(idx int) Node234 {
	var tempNode Node234
	tempNode = Node234{
		nil,
		nil,
		1,
		[]DataItem{
			{n.DataItems[idx].Data},
		},
		true,
	}
	n.DataItems[idx].Data = int(INF)
	return tempNode
}

//disconnect from node and return a new node that is to be a child at specified index
func (n *Node234) ConnectChild(idx int, node *Node234) error {
	n.Children[idx] = *node
	if &node != nil {
		node.Parent = n
	}
	 return nil
}

func (n *Node234) GetParent() (*Node234, error){
	parent := n.Parent
	if parent == nil {
		return nil, fmt.Errorf("parent is nil")
	}
	return parent, nil
}


func (n *Node234) GetNumItems() (int, error){
	return len(n.DataItems), nil
}

func (n *Node234) IsRoot() bool{
	return n.Parent == nil
}