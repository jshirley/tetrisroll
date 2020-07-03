package tetrisroll

import (
	"bytes"
	"crypto/sha256"
	"math/rand"
	"sort"
	"time"
)

type GlueSetItem struct {
	Key   string
	Count int
}

type GlueSet struct {
	items []*GlueSetItem
	key   []byte
	count int
}

func NewGlueSet(keys []string) *GlueSet {
	items := make([]*GlueSetItem, len(keys))
	var persistentKey bytes.Buffer

	sort.Strings(keys)
	for index, key := range keys {
		items[index] = &GlueSetItem{
			Key:   key,
			Count: 0,
		}
		persistentKey.WriteString(key)
	}

	key := sha256.Sum256([]byte(persistentKey.String()))

	rand.Seed(time.Now().UnixNano())

	return &GlueSet{
		key:   key[:],
		count: 0,
		items: items,
	}
}

func (gs *GlueSet) Key() []byte {
	return gs.key
}

func (gs *GlueSet) Count() int {
	return gs.count
}

func (gs *GlueSet) Roll() string {
	available := []*GlueSetItem{}

	for _, item := range gs.items {
		if item.Count == gs.count {
			available = append(available, item)
		}
	}
	// We didn't find anything, need to roll again!
	if len(available) == 0 {
		gs.count++
		return gs.Roll()
	}

	picked := rand.Intn(len(available))
	available[picked].Count += 1
	return available[picked].Key
}
