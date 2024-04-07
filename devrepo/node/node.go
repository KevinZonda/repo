package node

import (
	"encoding/json"
	"net/http"
)

const NODE_INDEX_URL = "https://nodejs.org/dist/index.json"

type Repo struct {
	hc *http.Client
}

func NewRepo() *Repo {
	return &Repo{hc: http.DefaultClient}
}

func (r *Repo) fetchIndex() (IndexJson, error) {
	resp, err := r.hc.Get(NODE_INDEX_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var index IndexJson
	err = json.NewDecoder(resp.Body).Decode(&index)
	return index, err
}
