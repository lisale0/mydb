package dsm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddDirectoryPage(t *testing.T) {
	dp := NewDirectoryPage()
	newEntry := NewFileEntry(1, "test.dat")
	dp.AddEntry(*newEntry)
	assert.Equal(t, dp.FileEntries[0].FirstPageId, PageId(1))
	assert.Equal(t, dp.FileEntries[0].FileName, "test.dat")
}
