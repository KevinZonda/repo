package node

import (
	"github.com/KevinZonda/repo/repo_standard"
	"log"
)

type NodeRepo struct {
	idx IndexJson
}

func (n *NodeRepo) GetPackage() repo_standard.Package {
	pkg := repo_standard.Package{
		DisplayName: "Node",
		UName:       "node",
		Category:    "Development",
	}
	if len(n.idx) == 0 {
		return pkg
	}
	latest := n.idx.Latest()
	lts := n.idx.LatestLts()
	pkg.Versions = repo_standard.VersionedUrls{
		"latest": latest.ToVersionedUrl(),
		"lts":    lts.ToVersionedUrl(),
	}
	return pkg
}

var r = NewRepo()

func (n *NodeRepo) Sync() {
	idx, err := r.fetchIndex()
	if err != nil {
		log.Println(err)
		log.Println("Failed to sync Node repo")
		return
	}
	n.idx = idx
}
