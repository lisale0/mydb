package formatter

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/lisale0/mydb/executor"
	"io"
)

const LatestVersion = 1

type Writer struct {
	writer      io.Writer
	uvarintBuf  []byte
	numRows     int
	columnNames []string
}

func NewWriter(columnNames []string, numRows int, w io.Writer) *Writer {
	return &Writer{
		writer:      w,
		uvarintBuf:  make([]byte, binary.MaxVarintLen64), //max 8 bytes
		numRows:     numRows,
		columnNames: columnNames,
	}
}

func (w *Writer) WriteData(data executor.Tuple) error {
	for _, i := range data.Values {
		_, err := w.writer.Write([]byte(i.StringValue))
		if err != nil {
			return fmt.Errorf("error writing to buffer")
		}
	}
	return nil
}

type Header struct {
	Version     int
	NumRows     int
	ColumnNames []string
}

func (w *Writer) WriteHeader() error {
	header := Header{
		Version:     LatestVersion,
		NumRows:     w.numRows,
		ColumnNames: w.columnNames,
	}
	headerBytes, err := json.Marshal(&header)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	/*length of header byte prepended before header data*/
	if err := w.writeUVarint(uint64(len(headerBytes))); err != nil {
		return fmt.Errorf("error writing the length of header %s", err)
	}
	if _, err = w.writer.Write(headerBytes); err != nil {
		return fmt.Errorf("error writeing header bytes %s", err)
	}
	return nil
}

func (w *Writer) WriteTuples(tuples []executor.Tuple) error {
	for _, tuple := range tuples {
		if err := w.WriteTuple(tuple.Values); err != nil {
			return fmt.Errorf("%v", err)
		}
	}
	return nil
}

func (w *Writer) WriteTuple(tupleValues []executor.Value) error {
	for _, value := range tupleValues {
		if err := w.writeUVarint(uint64(len(value.StringValue))); err != nil {
			return fmt.Errorf("error writing the length of tuple")
		}
		if _, err := w.writer.Write([]byte(value.StringValue)); err != nil {
			return fmt.Errorf("error writing tuple %s, err: %v", value.StringValue, err)
		}
	}
	return nil
}

func (w *Writer) writeUVarint(x uint64) error {
	varintLen := binary.PutUvarint(w.uvarintBuf, x)
	_, err := w.writer.Write(w.uvarintBuf[:varintLen])
	if err != nil {
		return fmt.Errorf("writeUVarint: error writing uvarint: %v", err)
	}
	return nil
}
