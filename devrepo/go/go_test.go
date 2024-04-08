package _go

import (
	"encoding/json"
	"github.com/KevinZonda/GoX/pkg/iox"
	"testing"
)

func TestNode(t *testing.T) {
	c := NewRepo()
	index, err := c.fetchIndex()
	if err != nil {
		t.Fatal(err)
		return
	}
	bs, _ := json.Marshal(index)
	iox.WriteAllBytes("go_index.json", bs)

	t.Log(index)
}

func TestGoLatest(t *testing.T) {
	bs, _ := iox.ReadAllByte("go_index.json")
	var index GoIndex
	json.Unmarshal(bs, &index)

	t.Log(index.Stable())

	t.Log(index.Latest())
}
