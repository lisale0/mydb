package formatter

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/lisale0/mydb/executor"
	"io"
)

// byteReader wraps an io.Reader so that it implements io.ByteReader.
type byteReader struct {
	io.Reader
	byteBuf []byte
}

// newByteReaders create a byteReader from an io.Reader.
func newByteReader(r io.Reader) *byteReader {
	return &byteReader{
		Reader:  r,
		byteBuf: make([]byte, 1),
	}
}

func (b *byteReader) ReadByte() (byte, error) {
	n, err := b.Reader.Read(b.byteBuf)
	if err != nil {
		return 0, fmt.Errorf("byteReader: ReadByte: error reading byte: %v", err)
	}
	if n != 1 {
		return 0, fmt.Errorf("byteReader: ReadByte: expected to read one byte, but read: %d", n)
	}

	return b.byteBuf[0], nil
}


type FileScanner struct {
	header  *Header
	r       *byteReader
	numRead int
	next    executor.Tuple
}


func NewFileScanner(r io.Reader) *FileScanner {
	return &FileScanner{
		r: newByteReader(r),
	}
}

func (f *FileScanner) ReadHeader(){
	/*grab the length of the header from the first byte*/
	headerLength, err := binary.ReadUvarint(f.r)
	if err != nil {
		fmt.Errorf("error getting the header length")
	}
	/*make a new byte buffer based on the headerlength*/
	headerBytes := make([]byte, headerLength)

	header := &Header{}

	/*read the following header bytes and then deserialize data into header obj*/
	if _, err := io.ReadFull(f.r, headerBytes); err != nil {
		panic(fmt.Sprintf("FileScanner: error reading header bytes: %v", err))
	}
	if err := json.Unmarshal(headerBytes, &header); err != nil {
		panic(fmt.Sprintf("FileScanner: error unmarshaling header: %v", err))
	}
	f.header = header
}

