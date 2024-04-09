package msys2

import (
	"github.com/KevinZonda/repo/repo_standard"
	"github.com/KevinZonda/repo/utils"
	"log"
)

type Msys2Repo struct {
	idx index
}

func (n *Msys2Repo) GetPackage() repo_standard.Package {
	pkg := repo_standard.Package{
		DisplayName: "MSYS2",
		UName:       "msys2",
		Category:    "Development",
	}
	if n.idx.winx64 == "" {
		return pkg
	}
	current := repo_standard.VersionedUrl{
		Version: utils.NotEmptyOne(n.idx.version, "stable"),
		Urls: map[repo_standard.Platform]map[repo_standard.Arch]string{
			repo_standard.PlatformWin: {
				repo_standard.ArchX64: utils.NotEmptyOne(n.idx.winx64),
			},
		},
	}
	pkg.Versions = repo_standard.VersionedUrls{
		"latest": current,
		"stable": current,
	}
	if n.idx.version != "" {
		pkg.Versions[n.idx.version] = current
	}
	pkg.History = pkg.Versions
	return pkg
}

func (n *Msys2Repo) Sync() {
	idx, err := fetchIndex()
	if err != nil {
		log.Println(err)
		log.Println("Failed to sync Git repo")
		return
	}
	n.idx = idx
}
