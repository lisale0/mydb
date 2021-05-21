package access

import (
	"fmt"
	"sort"
)

/*naive approach to B+-tree without concurrency implementation*/

const ROOTNODE = 1
const INTERNALNODE = 2
const LEAFNODE = 3


type BPlusTree struct {
	Root *BTreeNode
	Fanout int
}

type BTreeNode struct {
	Parent *BTreeNode
	NodeType int
	PrimaryKey int
	Keys [][]int
	Children []*BTreeNode

}

func NewBTreeNode(nodeType int, key int) BTreeNode{
	return BTreeNode{
		Parent: nil,
		NodeType: nodeType,
		PrimaryKey: key,
		Keys: [][]int{},
		Children: nil,
	}
}

func NewBPlusTree(fanout int) BPlusTree {
	return BPlusTree{
		Root: nil,
		Fanout: fanout,
	}
}

func (b *BPlusTree) search() error{
	return nil
}

func (b *BPlusTree) insert(key int) error{
	/*if the root is nil return a new root node*/
	if b.Root == nil {
		newRootNode := BTreeNode{nil,
			ROOTNODE,
			key,
			[][]int{{key}},
			nil,
		}
		b.Root = &newRootNode
		return nil
	}

	/*if there is still room in the keys for root*/
	if len(b.Root.Keys[0]) < b.Fanout-1{
		b.Root.Keys[0] = append(b.Root.Keys[0], key)
		sort.Ints(b.Root.Keys[0])
	} else {
		return b.split(*b.Root, key, 0)
	}
	return nil
}



func (b *BPlusTree) split(node BTreeNode, key int, idx int) error {
	lenKeys := len(node.Keys)
	keysInIdx := append(node.Keys[idx], key)
	sort.Ints(keysInIdx)
	splitLen := len(keysInIdx) / 2
	newParentKey := keysInIdx[splitLen]
	fmt.Print(newParentKey)


	leftSplit := keysInIdx[:splitLen]
	rightSplit := keysInIdx[splitLen:]

	newNode := BTreeNode{
		&node,
		INTERNALNODE,
		newParentKey,
		make([][]int, 4),
		nil,
	}

	newNode.Keys[idx] = leftSplit
	newNode.Keys[idx+1] = rightSplit
	//newNode.Children = append(node.Children, &newNode)


	/*update new root if new root*/
	if node.NodeType == ROOTNODE {
		newNode.NodeType = ROOTNODE
		b.Root = &newNode
		newNode.Parent = b.Root.Parent
	}
	newNode.Keys[0] = append(newNode.Keys[0], newParentKey)
	fmt.Print(lenKeys)
	return nil
}