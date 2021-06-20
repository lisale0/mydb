package buffermanager

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClock_New(t *testing.T) {
	clock := NewClock(HTSIZE)
	clock.PickVictim()

	clock.Entries = ClockEntry{
		123:  true,
		456:  true,
		768:  true,
		2313: true,
		3232: true,
		222:  true,
		111:  false,
		7656: true,
		3:    true,
		6:    true,
	}
	//returns the index of the entries that needs to be evicted
	pageToEvict := clock.PickVictim()
	assert.Equal(t, pageToEvict, 111)
}