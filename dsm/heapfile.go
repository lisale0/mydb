package dsm

type HeapFile struct {
	Name        string // name of the file
	FirstPage   PageId
	LastPage    PageId
	PageCount   int
	RecordCount int
}

func NewHeapFile(name string, firstPage PageId) *HeapFile {
	return &HeapFile{
		Name:        name,
		FirstPage:   firstPage,
		LastPage:    0,
		PageCount:   0,
		RecordCount: 0,
	}
}

func (h *HeapFile) GetRecordCount() {

}

func (h *HeapFile) InsertRecord() {

}

func (h *HeapFile) DeleteRecord() {

}

func (h *HeapFile) UpdateRecord() {

}

func (h *HeapFile) GetRecord() {

}

func (h *HeapFile) DeleteFile() {

}

type Scan struct {
}
