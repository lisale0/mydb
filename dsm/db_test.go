package dsm

import (
	"fmt"
	"github.com/lisale0/mydb/executor"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestAddFileEntry(t *testing.T) {
	db := NewDatabase("testdatabase")
	db.AddFileEntry("mydbtestmovies", 1)
	fmt.Print(db.Directory.FileEntries[0])
	assert.Equal(t, db.Directory.FileEntries[0].FileName, "mydbtestmovies")
	assert.Equal(t, db.Directory.FileEntries[0].FirstPageId, PageId(1))

	filePath := fmt.Sprintf("%s/%s", DATAPATH, "mydbtestmovies1")
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		t.Errorf("File failed to create: %s", "mydbtestmovies1")
	}
	os.Remove(filePath)
}

func TestWriteHeader(t *testing.T) {
	header := NewPageHeader(PageId(3))
	header.Columns = "id|gender"
	fileName := fmt.Sprintf("/tmp/testing%s.data", "32453243")
	f, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	bytes := []byte{
		0x03, 0x00, 0xff, 0xff, 0xff, 0xff, 0x00, 0x20, 0x00, 0x20, 0x00, 0x20, 0x00, 0x00, 0x09, 0x00,
		0x00, 0x00, 0x69, 0x64, 0x7c, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x00,
	}
	WriteHeader(*f, header)
	b, _ := ioutil.ReadFile(fileName)
	assert.Equal(t, b, bytes)
	f.Close()
}

func TestReadHeader(t *testing.T){
	bytes := []byte{
		0x03, 0x00, 0xff, 0xff, 0xff, 0xff, 0x00, 0x20, 0x00, 0x20, 0x00, 0x20, 0x00, 0x00, 0x09, 0x00,
		0x00, 0x00, 0x69, 0x64, 0x7c, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x00,
	}
	header, _ := ReadHeader(bytes)
	assert.Equal(t, header.PageId, PageId(3))
	assert.Equal(t, header.PrevPage, PageId(-1))
	assert.Equal(t, header.NextPage, PageId(-1))
	assert.Equal(t, header.FreeSpace, uint16(8192))
	assert.Equal(t, header.Upper, uint16(8192))
	assert.Equal(t, header.Lower, uint16(0))
	assert.Equal(t, header.ColumnSize, uint16(9))
	assert.Equal(t, header.Columns, "id|gender")
	assert.Equal(t, header.SlotCount, uint16(0))
}

func TestFlushRecord(t *testing.T) {
	var page Page
	header := NewPageHeader(PageId(3))
	header.Columns = "id|gender"

	page = NewPage(header)

	fileName := fmt.Sprintf("/tmp/testing%s.data", "32453243")
	f, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	bytes := []byte{
		0x08, 0x73, 0x74, 0x75, 0x64, 0x65,  0x6e, 0x74, 0x31, 0x04, 0x6d, 0x61, 0x6c, 0x65,
	}

	tuple := executor.NewTuple(
		"id", "student1",
		"gender", "male")
	page.insertRecord(tuple)

	WriteHeader(*f, page.PageHeader)
	WriteTuples(*f, page)
	b, _ := ioutil.ReadFile(fileName)
	//offset at 0x1ff0
	assert.Equal(t, bytes, b[len(b)-15: len(b)-1])
	f.Close()
}

func TestReadRecord(t *testing.T){
	fileName := fmt.Sprintf("/tmp/testing%s.data", "32453243")
	b, _ := ioutil.ReadFile(fileName)
	header, _ := ReadHeader(b)
	page := NewPage(header)
	ReadTuples(b, &page)
	assert.Equal(t, page.Records[0].Tuple.Values[0].StringValue, "student1")
	assert.Equal(t, page.Records[0].Tuple.Values[1].StringValue, "male")
}