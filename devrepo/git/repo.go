package git

import (
	"github.com/KevinZonda/repo/repo_standard"
	"github.com/KevinZonda/repo/utils"
	"log"
)

type GitRepo struct {
	idx index
}

func (n *GitRepo) GetPackage() repo_standard.Package {
	pkg := repo_standard.Package{
		DisplayName: "Git",
		UName:       "git",
		Category:    "Development",
	}
	if n.idx.winx64 == "" && n.idx.winx86 == "" {
		return pkg
	}
	current := repo_standard.VersionedUrl{
		Version: utils.NotEmptyOne(n.idx.version, "stable"),
		Urls: map[repo_standard.Platform]map[repo_standard.Arch]string{
			repo_standard.PlatformWin: {
				repo_standard.ArchX86: n.idx.winx86,
				repo_standard.ArchX64: utils.NotEmptyOne(n.idx.winx64, n.idx.winx86),
			},
		},
	}
	pkg.Versions = repo_standard.VersionedUrls{
		"latest": current,
		"stable": current,
	}
	pkg.History = pkg.Versions
	return pkg
}

func (n *GitRepo) Sync() {
	idx, err := fetchIndex()
	if err != nil {
		log.Println(err)
		log.Println("Failed to sync Git repo")
		return
	}
	n.idx = idx
}
