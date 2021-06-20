package access

import (
	"fmt"
	"github.com/lisale0/mydb/dsm"
	"github.com/lisale0/mydb/executor"
	"testing"
)

func Test_Indexing(t *testing.T) {
	header := dsm.NewPageHeader(dsm.PageId(5))
	newPage := dsm.NewPage(header)
	newPage.PageHeader.Columns = "name|city"
	tuple1 := executor.NewTuple("name", "Matt", "city", "NY")
	tuple2 := executor.NewTuple("name", "Dave", "city", "SF")
	tuple3 := executor.NewTuple("name", "Eve", "city", "Houston")
	tuple4 := executor.NewTuple("name", "Andy", "city", "Miami")
	tuple5 := executor.NewTuple("name", "Emily", "city", "Vancouver")
	tuple6 := executor.NewTuple("name", "Todd", "city", "Boise")
	tuple7 := executor.NewTuple("name", "Zack", "city", "Dubai")
	tuple8 := executor.NewTuple("name", "Bob", "city", "London")
	newPage.InsertRecord(tuple1)
	newPage.InsertRecord(tuple2)
	newPage.InsertRecord(tuple3)
	newPage.InsertRecord(tuple4)
	newPage.InsertRecord(tuple5)
	newPage.InsertRecord(tuple6)
	newPage.InsertRecord(tuple7)
	newPage.InsertRecord(tuple8)
	fmt.Print(newPage)
	NewIndex("")

}