package catalog

import (
	"errors"
	"github.com/lisale0/mydb/dsm"
)

type TableSpace struct {
	Name string
	SpaceLocation string
}

func NewTableSpace(name string, loc string) *TableSpace {
	return &TableSpace{
		Name: name,
		SpaceLocation: loc,
	}
}

// PageId mapped to table space.
type TableSpaceCatalog struct {
	TableSpaces map[dsm.PageId]TableSpace
}

func NewTableSpaceCatalog() *TableSpaceCatalog{
	return &TableSpaceCatalog{TableSpaces: map[dsm.PageId]TableSpace{}}
}

func (t *TableSpaceCatalog) AddTableSpace(id dsm.PageId, ts TableSpace) error {
	t.TableSpaces[id] = ts
	return nil
}

func (t *TableSpaceCatalog) GetTableSpace(id dsm.PageId) (*TableSpace, error){
	val, exists := t.TableSpaces[id]
	if exists {
		return &val, nil
	}
	return nil, errors.New("tablespace does not exist with the provided id")
}