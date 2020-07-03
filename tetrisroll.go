package tetrisroll

import (
	"bytes"
	"crypto/sha256"
	"math/rand"
	"sort"
	"time"
)

type SetItem struct {
	Key   string `json:"key" bson:"key"`
	Count int    `json:"count" bson:"count"`
}

type TetrisRollSet struct {
	items []*SetItem `json:"items" bson:"items"`
	key   []byte     `json:"key" bson:"_id"`
	count int        `json:"count" bson:"count"`
}

func NewTetrisRoll(keys []string) *TetrisRollSet {
	items := make([]*SetItem, len(keys))
	var persistentKey bytes.Buffer

	sort.Strings(keys)
	for index, key := range keys {
		items[index] = &SetItem{
			Key:   key,
			Count: 0,
		}
		persistentKey.WriteString(key)
	}

	key := sha256.Sum256([]byte(persistentKey.String()))

	rand.Seed(time.Now().UnixNano())

	return &TetrisRollSet{
		key:   key[:],
		count: 0,
		items: items,
	}
}

func (tr *TetrisRollSet) Key() []byte {
	return tr.key
}

func (tr *TetrisRollSet) Count() int {
	return tr.count
}

func (tr *TetrisRollSet) Roll() string {
	available := []*SetItem{}

	for _, item := range tr.items {
		if item.Count == tr.count {
			available = append(available, item)
		}
	}
	// We didn't find anything, need to roll again!
	if len(available) == 0 {
		tr.count++
		return tr.Roll()
	}

	picked := rand.Intn(len(available))
	available[picked].Count += 1
	return available[picked].Key
}
