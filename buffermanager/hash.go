package buffermanager

// BufHashTable
//The BufHashTable keeps track of pages in the buffer pool

const HTSIZE = 4

type HashEntry map[int]BufHashTableEntry

type BufHashTableEntry struct {
	Next     *BufHashTableEntry
	PageNum  int
	FrameNum int
}

// NewHashEntry created a new hash entry
func NewHashEntry(pageNum int, frameNum int) BufHashTableEntry {
	return BufHashTableEntry{
		nil,
		pageNum,
		frameNum,
	}
}

type BufHashTable struct {
	BufHashTableEntries []BufHashTableEntry //contains an array of Hash Table entries
}

func NewBufHashTable() *BufHashTable {
	return &BufHashTable{
		BufHashTableEntries: make([]BufHashTableEntry, HTSIZE),
	}
}

// Hash use a hash function to find the bucket in the hash "directory" table
func (b *BufHashTable) Hash(pageNum int) int {
	htIndex := pageNum % HTSIZE
	return htIndex
}

func (b *BufHashTable) Insert(pageNum int) {
	var newHashEntry BufHashTableEntry
	htIndex := b.Hash(pageNum)
	newHashEntry.FrameNum = htIndex
	newHashEntry.PageNum = pageNum
	b.BufHashTableEntries[htIndex] = newHashEntry
}

func (b *BufHashTable) Lookup() {

}

func (b *BufHashTable) Remove() {

}

func (b *BufHashTable) Display() {

}
