package dsm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/lisale0/mydb/executor"
	"github.com/lisale0/mydb/util"
	"os"
	"strconv"
	"strings"
)

const BUFSIZE int = 4
const HTSIZE int = 4
var DATAPATH = util.GetEnv("DATAPATH", "/tmp")

type Database struct {
	name      string
	Directory *DirectoryPage
}

func NewDatabase(name string) *Database {
	return &Database{
		name:      name,
		Directory: NewDirectoryPage(),
	}
}

// AddFileEntry Adds a file entry to the header page(s).
func (d *Database) AddFileEntry(fileName string, startPageNum int) (*os.File, error) {
	newEntry := NewFileEntry(PageId(startPageNum), fileName)

	name := fmt.Sprintf("%s/%s%s", DATAPATH, fileName, strconv.Itoa(startPageNum))
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	d.Directory.FileEntries = append(d.Directory.FileEntries, *newEntry)
	return f, nil
}

func (d *Database) GetFileEntry(fileName string) (*os.File, error) {
	filePath := fmt.Sprintf("%s/%s", DATAPATH, fileName)
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (d *Database) DeleteFileEntry(fileName string) error {
	filePath := fmt.Sprintf("%s/%s", DATAPATH, fileName)
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}


func WriteHeader(f os.File, header PageHeader) error {
	var err error
	col := []byte(header.Columns)
	fmt.Print(col)


	bufSize := 19+(4*header.SlotCount) + uint16(len(col))
	b := bytes.NewBuffer(make([]byte, bufSize))
	packet := b.Bytes()
	binary.LittleEndian.PutUint16(packet[0:], uint16(header.PageId))
	binary.LittleEndian.PutUint16(packet[2:], uint16(header.PrevPage))
	binary.LittleEndian.PutUint16(packet[4:], uint16(header.NextPage))
	binary.LittleEndian.PutUint16(packet[6:], header.FreeSpace)
	binary.LittleEndian.PutUint16(packet[8:],header.Size)
	binary.LittleEndian.PutUint16(packet[10:], header.Upper)
	binary.LittleEndian.PutUint16(packet[12:], header.Lower)
	binary.LittleEndian.PutUint16(packet[14:], uint16(len(col)))
	binary.LittleEndian.PutUint16(packet[16:], header.SlotCount)

	packetIndex := 18
	for i := 0; i < len(col); i++ {
		binary.LittleEndian.PutUint16(packet[packetIndex:], uint16(col[i]))
		packetIndex += 1
	}

	for i := 0; i < int(header.SlotCount); i++ {
		binary.LittleEndian.PutUint16(packet[packetIndex:], header.SlotArr[i].Offset)
		binary.LittleEndian.PutUint16(packet[packetIndex+2:], header.SlotArr[i].Length)
		packetIndex += 4
	}
	_, err = f.Write(packet)
	if err != nil {
		return err
	}

	return nil
}

func WriteTuples(f os.File, page Page) error {
	for i := 0; i < int(page.PageHeader.SlotCount); i++ {
		b := bytes.NewBuffer(make([]byte, page.PageHeader.SlotArr[i].Length + 1))
		createValueBuffer(*b, page.Records[i])
		f.WriteAt(b.Bytes(), int64(page.PageHeader.SlotArr[i].Offset))

	}
	return nil
}

func ReadTuples(b []byte, page *Page) error {
	var records []Record

	for i, slot := range page.PageHeader.SlotArr {
		startIdx := slot.Offset
		endIdx := slot.Offset + slot.Length
		tuple := parseTuple(b[startIdx: endIdx], page.PageHeader.Columns)
		record := NewRecord(page.PageHeader.PageId, int16(i), &tuple)
		records = append(records, *record)
	}
	page.Records = records
	return nil
}

// parseTuple extracts serialized a byte array into a tuple
// takes in param b to deserialize, and build tuple based on the column names passed in
func parseTuple(b []byte, columnName string) executor.Tuple {
	i := 0
	var values []executor.Value

	offset := 0
	columnArr := strings.Split(columnName, "|")
	for i < len(columnArr) {
		valLen := int(b[offset])
		offset += 1
		val := executor.Value{
			Name:        columnArr[i],
			StringValue: string(b[offset:valLen + offset]),
		}
		offset += valLen
		values = append(values, val)
		i++
	}

	return executor.NewTupleWithVal(values)
}
// probably a better way of doing this
func createValueBuffer(buf bytes.Buffer, record Record){
	var strLen int

	bufIndex := 0
	b := buf.Bytes()
	for _, value := range record.Tuple.Values {
		strLen = len(value.StringValue)
		binary.LittleEndian.PutUint16(b[bufIndex:], uint16(strLen))
		for i := 0; i < strLen; i++ {
			bufIndex += 1
			binary.LittleEndian.PutUint16(b[bufIndex:], uint16(value.StringValue[i]))
		}
		bufIndex += 1
	}
}

func ReadHeader(b []byte) (PageHeader, error){
	var pageHeader PageHeader

	buf := bytes.NewBuffer(b)

	binary.Read(buf, binary.LittleEndian, &pageHeader.PageId)
	binary.Read(buf, binary.LittleEndian, &pageHeader.PrevPage)
	binary.Read(buf, binary.LittleEndian, &pageHeader.NextPage)
	binary.Read(buf, binary.LittleEndian, &pageHeader.FreeSpace)
	binary.Read(buf, binary.LittleEndian, &pageHeader.Size)
	binary.Read(buf, binary.LittleEndian, &pageHeader.Upper)
	binary.Read(buf, binary.LittleEndian, &pageHeader.Lower)
	binary.Read(buf, binary.LittleEndian, &pageHeader.ColumnSize)
	binary.Read(buf, binary.LittleEndian, &pageHeader.SlotCount)

	colName := make([]byte, pageHeader.ColumnSize)
	for i := 0; i < int(pageHeader.ColumnSize); i++ {
		byteAtIndex, _ := buf.ReadByte()
		colName[i] = byteAtIndex
	}

	pageHeader.Columns = string(colName)
	pageHeader.SlotArr = make([]Slot, pageHeader.SlotCount)

	for i := 0; i < int(pageHeader.SlotCount); i++ {
		binary.Read(buf, binary.LittleEndian, &pageHeader.SlotArr[i].Offset)
		binary.Read(buf, binary.LittleEndian, &pageHeader.SlotArr[i].Length)
	}
	return pageHeader, nil
}