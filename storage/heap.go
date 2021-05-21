package storage

/*
* Slotted page implementation
*/

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/lisale0/mydb/executor"
	"github.com/lisale0/mydb/formatter"
	"io"
	"io/ioutil"
	"os"
	"unsafe"
)

type PageID int


// Page Directory -----------------------------------------------------------
type PageDirectoryItem struct {
	PageId uint8
	Location string  //Location of file
}

type PageDirectory struct {
	PageItem []PageDirectoryItem
}


// Block ---------------------------------------------------------------------
type Block struct {
	BlockID uint16
}

func NewBlock(blockID uint16) Block {
	return Block{
		BlockID: blockID,
	}
}

// Page Header-----------------------------------------------------------------
type LinePointer struct{
	Offset uint16
	Length uint16

}
type PageHeader struct {
	PageId uint16
	Upper uint16 //offset of start of free space
	Lower uint16 //offset to end of free space
	Size uint16
	Special uint16 //offset to start of special space, index access specific data
	NumRecs uint16 // number of records currently in file
	ItemIDData []LinePointer
}

func NewPageHeader(pageId uint16) PageHeader {
	return PageHeader{
		PageId: pageId,
		Upper: 8192,
		Lower: 2,
		Size: 8192,
		Special: 0,
		NumRecs: 0,
		ItemIDData:nil,
	}
}


func (p *PageHeader) WriteNewHeader(file string, pageId uint16) {
	f, _ := os.Create(file)
	f.Truncate(8192)
	block := NewBlock(1)
	pageheader := NewPageHeader(pageId)
	b := bytes.NewBuffer(make([]byte, 20))
    packet := b.Bytes()
    binary.LittleEndian.PutUint16(packet[0:], block.BlockID)
    binary.LittleEndian.PutUint16(packet[2:], pageheader.PageId)
    binary.LittleEndian.PutUint16(packet[4:], pageheader.Upper)
    binary.LittleEndian.PutUint16(packet[6:], pageheader.Lower)
    binary.LittleEndian.PutUint16(packet[8:], pageheader.Size)
    binary.LittleEndian.PutUint16(packet[10:], pageheader.Special)
    binary.LittleEndian.PutUint16(packet[12:], pageheader.NumRecs)
	f.WriteAt(packet, 0)

}
// Page ---------------------------------------------------------------------
type Page struct {
	Block Block
	PageHeader PageHeader
	Tuples []executor.Tuple
}

func NewPage(block Block, pageHeader PageHeader) Page {
	return Page{
		Block: block,
		PageHeader: pageHeader,
	}
}

func (p *Page) InsertTuple() {
	return
}

func (p *Page) LoadPage(file string, permissions int) (os.File, error) {
	f, err := os.OpenFile(file, permissions, 0755)
	if err != nil {
		return *f, err
	}
	b, _ := ioutil.ReadFile("testpage.dat")
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.LittleEndian, &p.Block)
	binary.Read(buf, binary.LittleEndian, &p.PageHeader.PageId)
	binary.Read(buf, binary.LittleEndian, &p.PageHeader.Upper)
	binary.Read(buf, binary.LittleEndian, &p.PageHeader.Lower)
	binary.Read(buf, binary.LittleEndian, &p.PageHeader.Size)
	binary.Read(buf, binary.LittleEndian, &p.PageHeader.Special)
	binary.Read(buf, binary.LittleEndian, &p.PageHeader.NumRecs)

	p.PageHeader.ItemIDData = make([]LinePointer, 2)

	for i := 0; i < int(p.PageHeader.NumRecs); i++ {
		binary.Read(buf, binary.LittleEndian, &p.PageHeader.ItemIDData[i].Offset)
		binary.Read(buf, binary.LittleEndian, &p.PageHeader.ItemIDData[i].Length)
	}
	return *f, nil
}


func (p *Page) WriteTuple(tuple executor.Tuple, f os.File) error {
	columnNames := []string{"id", "gender"}

	// write the tuple to the correct position
	buf := bytes.NewBuffer(nil)
	writer := formatter.NewWriter(columnNames, 1024, buf)
	writer.WriteTuple(tuple.Values)
	offset := p.PageHeader.Upper - uint16(unsafe.Sizeof(buf.Bytes()))
	f.WriteAt(buf.Bytes(), int64(offset))
	p.PageHeader.Upper = offset

	//add or update ItemIDData
	p.PageHeader.ItemIDData = append(p.PageHeader.ItemIDData, LinePointer{offset,uint16(len(buf.Bytes()))})
	p.PageHeader.NumRecs += 1
	return nil
}

func (p *Page) ReadTuples(){
	b, _ := ioutil.ReadFile("testpage.dat")

	buf := bytes.NewBuffer(b)
	for i := 0; i < int(p.PageHeader.NumRecs); i++ {
		// TODO get column metadata in order to get the number of values read
		offset := uint16(p.PageHeader.ItemIDData[i].Offset)
		length := uint16(buf.Bytes()[offset: offset + 1][0])
		data := buf.Bytes()[offset + 1: offset + 1 + length]
		fmt.Printf("%s",data)
	}

}

func (p *Page) FlushHeader(f io.Writer){
	b := bytes.NewBuffer(make([]byte, 20 + (4 * len(p.PageHeader.ItemIDData))))
    packet := b.Bytes()
    binary.LittleEndian.PutUint16(packet[0:], uint16(p.Block.BlockID))
    binary.LittleEndian.PutUint16(packet[2:], uint16(p.PageHeader.PageId))
    binary.LittleEndian.PutUint16(packet[4:], uint16(p.PageHeader.Upper))
    binary.LittleEndian.PutUint16(packet[6:], uint16(p.PageHeader.Lower))
    binary.LittleEndian.PutUint16(packet[8:], uint16(p.PageHeader.Size))
    binary.LittleEndian.PutUint16(packet[10:], uint16(p.PageHeader.Special))
    binary.LittleEndian.PutUint16(packet[12:], uint16(p.PageHeader.NumRecs))

    packetIndex := 14
    for i := 0; i < len(p.PageHeader.ItemIDData); i++ {
    	binary.LittleEndian.PutUint16(packet[packetIndex:], uint16(p.PageHeader.ItemIDData[i].Offset))
    	binary.LittleEndian.PutUint16(packet[packetIndex + 2:], uint16(p.PageHeader.ItemIDData[i].Length))
    	packetIndex += 4
	}
	fmt.Printf("%#v\n", packet)
    f.Write(packet)
}


/* look for env var MYDB_DATA and find where the data directory is */

func (p *Page) SetFileEntry(fileName string)(*os.File, error){
	/*
	filedir := os.Getenv("MYDB_DATA")
	pid := (p.PageHeader.PageId)
	fmt.Print(pid)
	name := filedir + "/" + fileName + pid
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	fmt.Print(f)
	*/

	return nil, nil
}