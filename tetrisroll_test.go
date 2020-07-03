package tetrisroll

import (
	"bytes"
	"sort"
	"testing"
)

func TestTetrisRoll(t *testing.T) {
	keys := []string{"Hello", "Item", "Another key", "One more", "🤔"}
	set1 := NewTetrisRoll(keys)
	set2 := NewTetrisRoll(keys)

	if bytes.Equal(set1.Key(), set2.Key()) != true {
		t.Errorf("NewTetrisRoll(%q).Key() != NewTetrisRoll(%q).Key() != %q", keys, keys, set1.Key())
	}

	for cycles := 0; cycles < 10; cycles++ {
		available := make([]string, len(keys))
		copy(available, keys)

		sort.Strings(available)

		rolls := 0
		remaining := len(available)

		for ok := true; ok; ok = rolls < remaining {
			rolled := set1.Roll()
			rolls += 1
			if set1.Count() != cycles {
				t.Errorf(".Count() != %d (got %d)", cycles, set1.Count())
			}

			i := sort.SearchStrings(available, rolled)
			if i < len(keys) && available[i] == rolled {
				// Found key, remove it from the slice
				available = append(available[:i], available[i+1:]...)
			} else {
				t.Errorf(".Roll() returned %q, expected one of %q", rolled, keys)
			}
		}
	}
}
