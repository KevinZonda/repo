package node

import (
	"github.com/KevinZonda/repo/utils"
)

const NODE_INDEX_URL = "https://nodejs.org/dist/index.json"

func fetchIndex() (IndexJson, error) {
	index, err := utils.HttpGetJson[IndexJson](NODE_INDEX_URL)
	return index, err
}
