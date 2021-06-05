package dsm

import (
	"github.com/lisale0/mydb/executor"
)

type Status int
type PageId int16

// RID Record Id
type RID struct {
	PageId  PageId
	SlotNum int16
}


type Slot struct {
	Offset uint16
	Length uint16
}

type Record struct {
	RecordId RID
	Tuple    executor.Tuple
}

type PageHeader struct {
	PageId    PageId
	PrevPage  PageId
	NextPage  PageId
	FreeSpace uint16
	Size      uint16
	Upper     uint16
	Lower     uint16
	ColumnSize uint16
	SlotCount uint16
	Columns string
	SlotArr   []Slot
}

func NewRecord(pageId PageId, slotNum int16, tuple *executor.Tuple) *Record{
	return &Record{
		RecordId: RID{
			pageId,
			slotNum,
		},
		Tuple:   *tuple,
	}
}

func NewSlot(offset uint16, length uint16) *Slot{
	return &Slot{
		Offset: offset,
		Length: length,
	}
}

func NewPageHeader(pageId PageId) PageHeader {
	return PageHeader{
		PageId: pageId,
		SlotArr:   []Slot{},
		SlotCount: 0,
		PrevPage:  -1,
		NextPage:  -1,
		FreeSpace: 8192,
		Size:      8192,
		Upper:     8192,
		Lower:     0,
		ColumnSize: 0,
		Columns: "",
	}
}

type Page struct {
	PageHeader PageHeader
	Records    []Record
}

func NewPage(pageHeader PageHeader) Page {
	return Page{
		PageHeader: pageHeader,
		Records: []Record{},
	}
}

func (p *Page) PageNum() PageId {
	return p.PageHeader.PageId
}

// set the next page for provided pageNum
// return: the next page id
func (p *Page) setNextPage(pageNum PageId) PageId {
	p.PageHeader.NextPage = pageNum
	return pageNum
}

// set the prev page for provided pageNum
// return: the prev pageid
func (p *Page) setPrevPage(pageNum PageId) PageId {
	p.PageHeader.PrevPage = pageNum
	return pageNum
}

// Get the next page with the provided pageId
func (p *Page) getNextPage() PageId {
	return p.PageHeader.NextPage
}

// Get the prev page with the provided pageId
func (p *Page) getPrevPage() PageId {
	return p.PageHeader.PrevPage
}

// inserts a new record pointed to by recPtr with length recLen onto
// the page, returns RID of record
func (p *Page) insertRecord(tuple executor.Tuple) Status {
	var newSlot Slot

	newRecord := NewRecord(p.PageHeader.PageId, int16(len(p.PageHeader.SlotArr)+1), &tuple)
	p.Records = append(p.Records, *newRecord)

	// Update slot arr
	valueLength := getRecordLength(p.Records[p.PageHeader.SlotCount])
	newSlot.Offset = p.PageHeader.Upper - uint16(valueLength)
	newSlot.Length = uint16(valueLength)
	p.PageHeader.SlotArr = append(p.PageHeader.SlotArr, newSlot)

	// increment slot count
	p.PageHeader.SlotCount += 1
	return 1
}

func getRecordLength(record Record) int{
	totalLength := len(record.Tuple.Values)

	for _, value := range record.Tuple.Values {
		totalLength += len(value.StringValue)
	}
	return totalLength
}

// markForDeletion will mark the record for deletion, actual deletion handled by vacuum (todo)
func (p *Page) markForDeletion(rid RID) Status {
	for _, record := range p.Records {
		if rid ==  record.RecordId {
			rid.PageId = -1
			rid.SlotNum = -1
		}
	}
	return 1
}

// delete the record with the specified rid
func (p *Page) deleteRecord(rid RID) Status {
	for i, record := range p.Records {
		if rid ==  record.RecordId {
			copy(p.Records[i:], p.Records[i+1:])
			p.Records[len(p.Records)-1] = Record{}
			p.Records = p.Records[:len(p.Records)-1]
		}
	}
	return 1
}

func (p *Page) firstRecord(curr RID, next RID) *Record {
	if len(p.Records) > 0 {
		return &p.Records[0]
	}
	return nil
}

// returns RID of next record on the page
// returns DONE if no more records exist on the page
func (p *Page) nextRecord(curr RID, next RID) (*Record, error) {
	var nextRec *Record
	var err error

	for i, record := range p.Records {
		if record.RecordId.PageId == curr.PageId &&
			record.RecordId.SlotNum == curr.SlotNum {
			nextRec, err = p.returnRecord(p.Records[i + 1].RecordId, i + 1)
			if err != nil {
				return nil, err
			}
			break
		}
	}
	return nextRec, nil
}

func (p *Page) returnRecord(rid RID, idx int) (*Record, error) {
	if idx != -1 {
		return &p.Records[idx], nil
	}
	for i, _ := range p.Records{
		if p.Records[i].RecordId == rid {
			return &p.Records[i], nil
		}
	}
	return nil, nil
}
