package buffermanager


type ReplacementState int
type BufErrorCodes int

const (
	HASH_TBL_ERROR BufErrorCodes = iota
	HASH_NOT_FOUND
	BUFFER_EXCEEDED
	PAGE_NOT_PINNED
	BAD_BUFFER
	PAGE_PINNED
	REPLACER_ERROR
	BAD_BUF_FRAMENO
	PAGE_NOT_FOUND
	FRAME_EMPTY
)

const (
	AVAILABLE ReplacementState = iota
	REFERENCE
	PINNED
)

type Replacer interface {
	PickVictim() int
}

type ClockEntry map[int]bool

type Clock struct {
	Name    string
	Entries ClockEntry
	Size    int
}

func NewClock(size int) *Clock {
	return &Clock{
		Name:    "clock",
		Entries: ClockEntry{},
		Size:    size,
	}
}

func (c *Clock) PickVictim() int {
	var pageNumToEvict int
	for k, v := range c.Entries {
		if v == true {
			c.Entries[k] = false
		} else {
			pageNumToEvict = k
			break
		}
	}
	return pageNumToEvict
}