package buffermanager

import (
	"github.com/lisale0/mydb/catalog"
	"github.com/lisale0/mydb/dsm"
	"github.com/lisale0/mydb/executor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBufMgr_PinPage(t *testing.T) {
	fileName := "/tmp/db-test-student.dat"

	tsCatalog = catalog.NewTableSpaceCatalog()
	tsCatalog.AddTableSpace(dsm.PageId(5), *catalog.NewTableSpace("students", fileName))
	header := dsm.NewPageHeader(dsm.PageId(5))
	newPage := dsm.NewPage(header)
	newPage.PageHeader.Columns = "id|gender"
	tuple := executor.NewTuple("id", "student54", "gender", "other")
	newPage.InsertRecord(tuple)
	dsm.WritePage(fileName, newPage)


	/*setup*/
	clock := NewClock(4)
	clock.Entries = ClockEntry{
		123: true,
		456: true,
		768: true,
		21:  false,      //this is the one to evict if we need to replace
	}
	frameDescTable := []FrameDesc{
		NewFrameDesc(123),
		NewFrameDesc(456),
		NewFrameDesc(768),
		NewFrameDesc(21),
	}
	frameDescTable[0].PinCount = 1
	frameDescTable[2].PinCount = 1
	bufHashTable := BufHashTable{
		[]BufHashTableEntry{
			NewHashEntry(123, 0),
			NewHashEntry(456, 1),
			NewHashEntry(768, 2),
			NewHashEntry(21, 3),
		},
	}

	bufPool := []dsm.Page{
		dsm.NewPage(dsm.NewPageHeader(123)),
		dsm.NewPage(dsm.NewPageHeader(456)),
		dsm.NewPage(dsm.NewPageHeader(768)),
		dsm.NewPage(dsm.NewPageHeader(21)),
	}

	bufferManager := NewBufMgr(4, clock, frameDescTable)
	bufferManager.BufferPool = bufPool
	bufferManager.BufHashTable = bufHashTable
	bufferManager.PinPage(5, 0)
	assert.Equal(t, bufPool[3].PageHeader.PageId, dsm.PageId(5))
	assert.Equal(t, len(bufPool[3].Records), 1)
}


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
