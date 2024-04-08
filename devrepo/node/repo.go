package node

import "github.com/KevinZonda/repo/repo_standard"

type NodeRepo struct {
	idx IndexJson
}

func (n NodeRepo) GetPackage() repo_standard.Package {
	if len(n.idx) == 0 {
		return repo_standard.Package{}
	}
	latest := n.idx.Latest()
	lts := n.idx.LatestLts()
	return repo_standard.Package{
		DisplayName: "Node",
		UName:       "node",
		Category:    "Development",
		Urls: repo_standard.VersionedUrls{
			"latest": latest.ToVersionedUrl(),
			"lts":    lts.ToVersionedUrl(),
		},
	}
}

var r = NewRepo()

func (n *NodeRepo) Sync() {
	n.idx, _ = r.fetchIndex()
}
