package dsm

import (
	"github.com/lisale0/mydb/executor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPage(t *testing.T) {
	pageHeader := NewPageHeader(2)
	page := NewPage(pageHeader)
	assert.Equal(t, page.PageHeader.PageId, PageId(2))
}

func TestInsertRecord(t *testing.T) {
	pageHeader := NewPageHeader(2)
	page := NewPage(pageHeader)
	tuple := executor.NewTuple(
		"id", "student1",
		"gender", "male")
	page.insertRecord(tuple)
	assert.Equal(t, page.Records[0].Tuple.Values[0].StringValue, "student1")
	assert.Equal(t, page.Records[0].Tuple.Values[1].StringValue, "male")
}

