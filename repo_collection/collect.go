package repo_collection

import (
	"github.com/KevinZonda/repo/devrepo/node"
	"github.com/KevinZonda/repo/entertainrepo"
	"github.com/KevinZonda/repo/repo_standard"
)

var R = []repo_standard.IRepo{
	entertainrepo.Steam{},
	entertainrepo.Telegram{},
	node.NodeRepo{},
}
