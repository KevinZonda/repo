package _go

import "github.com/KevinZonda/repo/repo_standard"

type GoRepo struct {
	idx GoIndex
}

func (n GoRepo) GetPackage() repo_standard.Package {
	if len(n.idx) == 0 {
		return repo_standard.Package{}
	}
	latest := n.idx.Latest()
	stable := n.idx.Stable()
	return repo_standard.Package{
		DisplayName: "Go",
		UName:       "go",
		Category:    "Development",
		Urls: repo_standard.VersionedUrls{
			"latest": latest.VersionedUrl(),
			"stable": stable.VersionedUrl(),
		},
	}
}

var r = NewRepo()

func (n *GoRepo) Sync() {
	n.idx, _ = r.fetchIndex()
}
