package node

import (
	"github.com/KevinZonda/repo/utils"
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
	index, err := utils.HttpGetJson[IndexJson](NODE_INDEX_URL)
	return index, err
}
