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
	latest := n.idx.Latest().VersionedUrl()
	stable := n.idx.Stable().VersionedUrl()
	pkg.Versions = repo_standard.VersionedUrls{
		"latest": latest,
		"stable": stable,
	}
	pkg.History = n.allVersions()
	pkg.History["latest"] = latest
	pkg.History["stable"] = stable

	return pkg
}

func (n *GoRepo) allVersions() map[string]repo_standard.VersionedUrl {
	var m = map[string]repo_standard.VersionedUrl{}
	for _, v := range n.idx {
		url := v.VersionedUrl()
		m[v.Version] = url
	}
	return m
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
