package dsm

/*
The DB class also provides a file naming service, which is used by higher-level code to create logical ``files of pages''. This service is implemented using records consisting of file names and their header page ids. There are functions to insert, look up, and delete file entries. The set of file entries is collectively referred to as the file directory.
}
*/

type FileEntry struct {
	FirstPageId PageId
	FileName    string
}

func NewFileEntry(pageId PageId, fileName string) *FileEntry {
	return &FileEntry{
		FirstPageId: pageId,
		FileName:    fileName,
	}
}

type DirectoryPage struct {
	FileEntries []FileEntry
}

func NewDirectoryPage() *DirectoryPage {
	return &DirectoryPage{
		FileEntries: []FileEntry{},
	}
}

func (d *DirectoryPage) AddEntry(entry FileEntry) {
	d.FileEntries = append(d.FileEntries, entry)
}
