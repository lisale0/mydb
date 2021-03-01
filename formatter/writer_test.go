package formatter

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/lisale0/mydb/executor"
	"testing"
)

func TestWriteHeader(t *testing.T) {
	columnNames := []string{"id", "gender"}
	buf := bytes.NewBuffer(nil)

	writer := NewWriter(columnNames, 1024, buf)
	writer.WriteHeader()
}

func TestWriteTuple(t *testing.T) {
	columnNames := []string{"id", "gender"}
	buf := bytes.NewBuffer(nil)
	writer := NewWriter(columnNames, 1024, buf)
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
	writer.WriteHeader()
	writer.WriteTuples(tuples)
	fmt.Print(hex.Dump(buf.Bytes()))
}