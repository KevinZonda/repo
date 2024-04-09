package node

import (
	"encoding/json"
	"github.com/KevinZonda/GoX/pkg/iox"
	"testing"
)

func TestNode(t *testing.T) {
	index, err := fetchIndex()
	if err != nil {
		t.Fatal(err)
		return
	}
	bs, _ := json.Marshal(index)
	iox.WriteAllBytes("node_index.json", bs)

	t.Log(index)
}

func TestNodeLatest(t *testing.T) {
	bs, _ := iox.ReadAllByte("node_index.json")
	var index IndexJson
	json.Unmarshal(bs, &index)

	t.Log(index.LatestLts().DownloadList())

	t.Log(index.Latest().DownloadList())
}
