package _go

import (
	"github.com/KevinZonda/repo/repo_standard"
	"log"
)

type GoRepo struct {
	idx GoIndex
}

func (n *GoRepo) GetPackage() repo_standard.Package {
	pkg := repo_standard.Package{
		DisplayName: "Go",
		UName:       "go",
		Category:    "Development",
	}
	if len(n.idx) == 0 {
		return pkg
	}
	latest := n.idx.Latest()
	stable := n.idx.Stable()
	pkg.Versions = repo_standard.VersionedUrls{
		"latest": latest.VersionedUrl(),
		"stable": stable.VersionedUrl(),
	}
	return pkg
}

var r = NewRepo()

func (n *GoRepo) Sync() {
	idx, err := r.fetchIndex()
	if err != nil {
		log.Println(err)
		log.Println("Failed to sync go repo")
		return
	}
	n.idx = idx
}
