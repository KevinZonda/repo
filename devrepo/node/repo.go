package node

import (
	"github.com/KevinZonda/repo/repo_standard"
	"log"
	"strings"
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
	latest := n.idx.Latest().ToVersionedUrl()
	lts := n.idx.LatestLts().ToVersionedUrl()
	pkg.Versions = repo_standard.VersionedUrls{
		"latest": latest,
		"lts":    lts,
		"stable": lts,
	}
	pkg.History = n.allVersions()
	pkg.History["latest"] = latest
	pkg.History["lts"] = lts
	pkg.History["stable"] = lts
	return pkg
}

func (n *NodeRepo) allVersions() map[string]repo_standard.VersionedUrl {
	var m = map[string]repo_standard.VersionedUrl{}

	for _, v := range n.idx {
		url := v.ToVersionedUrl()
		m[v.Version] = url
		m[strings.TrimLeft(v.Version, "v")] = url
	}
	return m

}

func (n *NodeRepo) Sync() {
	idx, err := fetchIndex()
	if err != nil {
		log.Println(err)
		log.Println("Failed to sync Node repo")
		return
	}
	n.idx = idx
}
