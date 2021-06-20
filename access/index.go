package access

import "github.com/lisale0/mydb/dsm"

type Index struct {
	Name string    // name of the index
	KeyName string // key in the table to index
	Pointer dsm.PageId
}

func NewIndex(keyName string, pageId dsm.PageId) *Index{
	return &Index{
		KeyName: keyName,
		Name: "",
		Pointer: pageId,
	}
}

func (*Index) IndexPage(page dsm.Page){

}