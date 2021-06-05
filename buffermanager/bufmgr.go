package buffermanager

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/lisale0/mydb/dsm"
	"github.com/lisale0/mydb/util"
	"io/ioutil"
	"os"
)

/*
The buffer pool is  a collection of frames (FrameDesc), managed by the buffer manager

A simple hash table is used to figure out what frame a given disk page occupies. (BufHashTable)
The hash table is implemented entirely in main memory based on <pageNumber,frameNumber>

The hash table contains an array called the directory
and each list pair is called a bucket

The hash function find the directory entry point to the bucket that contains the frame number for a given page


|--------------------------------------------------------------------------------------------------------|
|--------------------------------------------------------------------------------------------------------|
|											 BUFFER MANAGER                                              |
|--------------------------------------------------------------------------------------------------------|
|--------------------------------------------------------------------------------------------------------|                                                                                                     |
|   HashTable "aka Directory"                                                                            |
|  |---------|                                                                                           |
|  |bucket 1 |           bucket                       FrameTable            frame                        |
|  |---------|         |--------|                    |---------|          |---------|                    |
|  |bucket 2-----------|-> next |     |------------->| frame 1------->    | pageNum |   --------|        |
|  |---------|         |--------|     |              |---------|          |---------|   ------| |        |
|  |  ...    |         | pageNo |---- |              | frame 2 |          |  dirty  |         | |        |
|  |---------|         |--------|                    |---------|          |---------|         | |        |
|  |bucket N |         |frameNo |                    |  . . .  |          | pincount|         | |        |
|  |---------|         |--------|                    |---------|          |---------|         | |        |
|                                                    | frame N |                              | |        |
|                                                    |---------|                              | |        |
|       --------------------------------------------------------------------------------------  |        |
|       | |-------------------------------------------------------------------------------------|        |-------------
|        V                                                                                               |
|        V                                                                                               |
|                                                                                                        |
|   BufferPool
|  |---------|                                                                                           |
|  | page 1  |                                                                                           |
|  |---------|                                                                                           |
|  | page 2  |                                                                                           |
|  |---------|                                                                                           |
|  |  . . .  |                                                                                           |
|  |---------|                                                                                           |
|  | page N  |                                                                                           |
|  |---------|                                                                                           |
|                                                                                                        |
|--------------------------------------------------------------------------------------------------------|
|                                                                                                        |
| Replacement Inteface                                                                                   |
| |-------------------|                                                                                  |
| |  pickvictim()     |                                                                                  |
| |-------------------|                                                                                  |
|   * algorithm for eviction will vary depending                                                         |
|     on types of implementation (i.e LRU, Clock, etc)                                                   |
|                                                                                                        |
|                                                                                                        |
|--------------------------------------------------------------------------------------------------------|


The BufMgr checks the HashTable to see if a page exists, if it does, the
hash table is used to reference the bufferpool


*/


// FrameDesc -----------------------------------
type FrameDesc struct {
	PageNum  int
	Dirty    bool
	PinCount int
}

func NewFrameDesc(pageNum int) FrameDesc {
	return FrameDesc{
		PageNum:  pageNum,
		Dirty:    false,
		PinCount: 0,
	}
}

func (f *FrameDesc) Pin() {
	return
}

func (f *FrameDesc) Unpin() {
	return
}



// BufMgr  ----------------------------------------------
type BufMgr struct {
	BufferPool   []dsm.Page
	BufSize      int
	BufHashTable BufHashTable
	FrameTable   []FrameDesc
	Replacer     Replacer
}

func NewBufMgr(bufsize int, replacer Replacer, frameTable []FrameDesc) *BufMgr {
	return &BufMgr{
		BufferPool:   make([]dsm.Page, bufsize),
		BufHashTable: *NewBufHashTable(),
		BufSize:      bufsize,
		Replacer:     replacer,
		FrameTable:   frameTable,
	}
}

/**
 * Pin a page.
 * First check if this page is already in the buffer pool.
 * If it is, increment the pin_count and return a pointer to this
 * page.
 * If the pin_count was 0 before the call, the page was a
 * replacement candidate, but is no longer a candidate.
 * If the page is not in the pool, choose a frame (from the
 * set of replacement candidates) to hold this page, read the
 * page (using the appropriate method from {\em diskmgr} package) and pin it.
 * Also, must write out the old page in chosen frame if it is dirty
 * before reading new page.__ (You can assume that emptyPage==false for
 * this assignment.)
 *
 * @param pageno page number
 * @param emptyPage true (empty page); false (non-empty page)
 */

func (b *BufMgr) PinPage(pageNum int, emptyNum int) *dsm.Page {
	potentialCandidate := make(map[int]int, 4)
	/** check if page is in buffer pool
	 * if exists, increment pin count and return pointer to this page
	 */
	idx := b.BufHashTable.Hash(pageNum)
	if b.BufHashTable.BufHashTableEntries[idx].PageNum == pageNum {
		b.FrameTable[idx].PinCount += 1
		return &b.BufferPool[idx]
	}
	// write out old values in the page before eviction if dirty bit is true
	for i, v := range b.BufHashTable.BufHashTableEntries {
		fmt.Print(i)
		fmt.Print(v)
		frame := b.FrameTable[v.FrameNum]
		fmt.Print(frame)
		if frame.PinCount == 0 {
			potentialCandidate[i] = frame.PageNum
		}

	}

	evictPage := b.Replacer.PickVictim()
	evictIndex := b.BufHashTable.Hash(evictPage)

	if b.FrameTable[evictIndex].Dirty == true {
		fmt.Print("flush page before evicting")
	}
	for i, p := range b.BufferPool {
		if p.PageHeader.PageId == dsm.PageId(evictPage) {
			var page dsm.Page
			//TODO find the file name associated with the pageID
			//page.LoadPage("", 0755)
			b.FrameTable[evictIndex].PageNum = pageNum
			b.BufHashTable.BufHashTableEntries[evictIndex].PageNum = pageNum
			b.BufferPool[i] = page
			continue
		}
	}
	return nil
}

/**
 * Unpin a page specified by a pageId.
 * This method should be called with dirty==true if the client has
 * modified the page.
 * If so, this call should set the dirty bit
 * for this frame.
 * Further, if pin_count>0, this method should
 * decrement it.
 * If pin_count=0 before this call, throw an exception
 * to report error.
 * (For testing purposes, we ask you to throw
 * an exception named PageUnpinnedException in case of error.)
 *
 * @param pageno page number in the Minibase.
 * @param dirty the dirty bit of the frame
 */
func (b *BufMgr) UnpinPage(pageNum int, dirty bool) error {
	if dirty == true {
		idx := b.BufHashTable.Hash(pageNum)
		framePinCount := b.FrameTable[idx].PinCount
		if framePinCount > 0 {
			b.FrameTable[idx].PinCount -= 1
		} else {
			return errors.New("PageUnpinnedException")
		}
	}
	return nil
}

// TODO implement the methods below
/**
 * Allocate new pages.
 * Call DB object to allocate a run of new pages and
 * find a frame in the buffer pool for the first page
 * and pin it. (This call allows a client of the Buffer Manager
 * to allocate pages on disk.) If buffer is full, i.e., you
 * can't find a frame for the first page, ask DB to deallocate
 * all these pages, and return null.
 *
 * @param firstpage the address of the first page.
 * @param howmany total number of allocated new pages.
 *
 * @return the first page id of the new pages.__ null, if error.
 */
func (b *BufMgr) NewPage(firstPage dsm.Page, numOfPages int) {
	// use db object to allocate new pages

	// return first page id
}

func (b *BufMgr) FreePage(globalPageId int) {
	// delete page on disk
}

// flush a particular page
func (b *BufMgr) FlushPage(pageNum int) {
	// if dirty, then flush
}

// Flush all dirty pages
func (b *BufMgr) FlushAllPages() {

}

// get the total number of unpinned frames
func (b *BufMgr) GetUnpinnedNum() {

}

func (d *BufMgr) WritePage(pageName string, pageNum int, pagePtr dsm.Page) (*os.File, error) {
	var f *os.File
	var err error

	dataPath := util.GetEnv("DATAPATH", "/tmp")
	filePath := fmt.Sprintf("%s/%s", dataPath, pageName)
	if util.FileExists(filePath) {
		f, err = util.OpenFile(filePath)
	} else {
		f, err = util.CreateFile(filePath)
	}
	if err != nil {
		return nil, err
	}
	fmt.Print(f)
	return nil, nil
}

func (d *BufMgr) ReadPage(file string, permissions int, p *dsm.Page) (*os.File, error) {
	f, err := os.OpenFile(file, permissions, 0755)
	if err != nil {
		return f, err
	}
	b, _ := ioutil.ReadFile("testpage.dat")
	buf := bytes.NewBuffer(b)
	//binary.Read(buf, binary.LittleEndian, &p.Block)
	//binary.Read(buf, binary.LittleEndian, &p.PageHeader.PageId)
	//binary.Read(buf, binary.LittleEndian, &p.PageHeader.Upper)
	//binary.Read(buf, binary.LittleEndian, &p.PageHeader.Lower)
	//binary.Read(buf, binary.LittleEndian, &p.PageHeader.Size)
	//binary.Read(buf, binary.LittleEndian, &p.PageHeader.Special)
	//binary.Read(buf, binary.LittleEndian, &p.PageHeader.NumRecs)

	p.PageHeader.SlotArr = make([]dsm.Slot, 2)

	for i := 0; i < int(p.PageHeader.SlotCount); i++ {
		binary.Read(buf, binary.LittleEndian, &p.PageHeader.SlotArr[i].Offset)
		binary.Read(buf, binary.LittleEndian, &p.PageHeader.SlotArr[i].Length)
	}
	return f, nil
}