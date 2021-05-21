package storage

import (
	"encoding/binary"
	"github.com/lisale0/mydb/executor"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestPageCreation(t *testing.T){
	var page PageHeader
	page.WriteNewHeader("testpage.dat", 1)
    file, _ := ioutil.ReadFile("testpage.dat")
    assert.Equal(t, binary.LittleEndian.Uint16(file[8:]), uint16(8192))
}


func TestWriteTuple(t *testing.T){
	var page Page
	f, _ := page.LoadPage("testpage.dat", os.O_RDWR)
	tuples := []executor.Tuple{
		executor.NewTuple(
			"id", "student1",
			"gender", "male"),
		executor.NewTuple(
			"id", "student2",
			"gender", "female"),
		executor.NewTuple(
			"id", "student3",
			"gender", "female"),
	}
	page.WriteTuple( tuples[0], f)
	page.WriteTuple(tuples[1], f)
	page.FlushHeader(&f)
}

func TestFlushPageHeader(t *testing.T){
	var page Page
	f, _ := page.LoadPage("testpage.dat", os.O_RDWR)

	page.PageHeader.NumRecs = 9
	page.FlushHeader(&f)
}


func TestReadHeader(t *testing.T){
	var page Page
	f, _ := page.LoadPage("testpage.dat", os.O_RDWR)
	f.Truncate(8192)
}

func TestReadTuple(t *testing.T){
	var page Page
	f, _ := page.LoadPage("testpage.dat", os.O_RDWR)
	f.Truncate(8192)
	page.ReadTuples()
}
