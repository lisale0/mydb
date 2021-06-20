package catalog

import (
	"github.com/lisale0/mydb/dsm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTableSpace(t *testing.T) {
	pageId := dsm.PageId(9)
	catalog := NewTableSpaceCatalog()
	catalog.AddTableSpace(pageId, *NewTableSpace("Employees", "/tmp/employees.dat"))
	assert.Equal(t, catalog.TableSpaces[pageId].Name, "Employees")
	assert.Equal(t, catalog.TableSpaces[pageId].SpaceLocation, "/tmp/employees.dat")
}

func TestGetTableSpace(t *testing.T) {
	pageId := dsm.PageId(9)
	catalog := NewTableSpaceCatalog()
	catalog.TableSpaces[pageId] = TableSpace{
		Name:          "Employee",
		SpaceLocation: "/tmp/employees.dat",
	}
	ts, _ := catalog.GetTableSpace(pageId)
	assert.Equal(t, ts.Name, "Employee")
	assert.Equal(t, ts.SpaceLocation, "/tmp/employees.dat")
}

func TestGetTableSpace_Error(t *testing.T) {
	catalog := NewTableSpaceCatalog()
	pageId := dsm.PageId(9)
	_, err := catalog.GetTableSpace(pageId)
	assert.EqualError(t, err, "tablespace does not exist with the provided id")
}