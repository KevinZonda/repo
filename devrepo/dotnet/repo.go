package dotnet

import (
	"github.com/KevinZonda/repo/repo_standard"
	"log"
)

type DotNetRepo struct {
	idx idx
}

type idx struct {
	globalIdx  ReleaseIndexJson
	versionIdx map[string]DotnetRelease
}

func (n *DotNetRepo) GetPackage() repo_standard.Package {
	pkg := repo_standard.Package{
		DisplayName: ".NET",
		UName:       "dotnet",
		Category:    "Development",
	}
	return pkg
}

func (n *DotNetRepo) allVersions() map[string]repo_standard.VersionedUrl {
	var m = map[string]repo_standard.VersionedUrl{}
	return m

}

var r = DotNetRepo{}

func (d DotNetRepo) fetchIndex() (idx, error) {
	return d.idx, nil
}

func (n *DotNetRepo) Sync() {
	idx, err := r.fetchIndex()
	if err != nil {
		log.Println(err)
		log.Println("Failed to sync Node repo")
		return
	}
	n.idx = idx
}
