package buffermanager

import (
	"testing"
)



//
//func TestBufMgr_PinPage(t *testing.T) {
//	clock := NewClock(HTSIZE)
//	clock.Entries = ClockEntry{
//		123: true,
//		456: true,
//		768: true,
//		21:  false,
//	}
//	frameDescTable := []FrameDesc{
//		NewFrameDesc(768),
//		NewFrameDesc(21),
//		NewFrameDesc(646),
//		NewFrameDesc(123),
//	}
//
//	frameDescTable[0].PinCount = 1
//	frameDescTable[2].PinCount = 1
//
//	bufHashTable := BufHashTable{
//		[]BufHashTableEntry{
//			NewHashEntry(768, 0),
//			NewHashEntry(21, 1),
//			NewHashEntry(646, 2),
//			NewHashEntry(123, 3),
//		},
//	}
//	bufPool := []storage.Page{
//		storage.NewPage(storage.Block{}, storage.PageHeader{
//			768, 0, 0, 0, 0, 0, nil,
//		}),
//		storage.NewPage(storage.Block{}, storage.PageHeader{
//			21, 0, 0, 0, 0, 0, nil,
//		}),
//		storage.NewPage(storage.Block{}, storage.PageHeader{
//			646, 0, 0, 0, 0, 0, nil,
//		}),
//		storage.NewPage(storage.Block{}, storage.PageHeader{
//			123, 0, 0, 0, 0, 0, nil,
//		}),
//	}
//	bufmgr := NewBufMgr(4, clock, frameDescTable)
//	bufmgr.BufHashTable = bufHashTable
//	bufmgr.BufferPool = bufPool
//	bufmgr.PinPage(567, 0)
//	//assert.Equal(t, bufmgr.BufferPool[1].PageHeader.PageId, 567)
//	assert.Equal(t, bufmgr.BufHashTable.BufHashTableEntries[1].PageNum, 567)
//
//}

func TestBufMgr_UnPinPage(t *testing.T) {
	clock := NewClock(4)
	clock.Entries = ClockEntry{
		123: true,
		456: true,
		768: true,
		111: false,
	}
	frameDescTable := []FrameDesc{
		NewFrameDesc(768),
		NewFrameDesc(21),
		NewFrameDesc(646),
		NewFrameDesc(123),
	}

	frameDescTable[0].PinCount = 1
	frameDescTable[0].Dirty = true
	frameDescTable[2].PinCount = 1

	bufHashTable := BufHashTable{
		[]BufHashTableEntry{
			NewHashEntry(768, 0),
			NewHashEntry(21, 1),
			NewHashEntry(646, 2),
			NewHashEntry(123, 3),
		},
	}

	bufmgr := NewBufMgr(4, clock, frameDescTable)

	bufmgr.BufHashTable = bufHashTable
	bufmgr.UnpinPage(768, true)

}
