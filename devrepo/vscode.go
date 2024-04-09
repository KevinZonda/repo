package devrepo

import "github.com/KevinZonda/repo/repo_standard"

type VSCodeRepo struct {
	repo_standard.BaseSync
}

func (v VSCodeRepo) GetPackage() repo_standard.Package {
	m := map[repo_standard.Platform]map[repo_standard.Arch]string{
		repo_standard.PlatformWin: {
			repo_standard.ArchX64: "https://code.visualstudio.com/sha/download?build=stable&os=win32-x64-user",
		},
		repo_standard.PlatformMac: {
			repo_standard.ArchX64:   "https://code.visualstudio.com/sha/download?build=stable&os=darwin-universal",
			repo_standard.ArchARM64: "https://code.visualstudio.com/sha/download?build=stable&os=darwin-universal",
		},
	}
	return repo_standard.Package{
		DisplayName: "VSCode",
		UName:       "vscode",
		Category:    "Development",
		Versions: map[string]repo_standard.VersionedUrl{
			"stable": {
				Version: "stable",
				Urls:    m,
			},
			"latest": {
				Version: "latest",
				Urls:    m,
			},
			"insider": {
				Version: "insider",
				Urls: map[repo_standard.Platform]map[repo_standard.Arch]string{
					repo_standard.PlatformWin: {
						repo_standard.ArchX64: "https://code.visualstudio.com/sha/download?build=insider&os=win32-x64-user",
					},
					repo_standard.PlatformMac: {
						repo_standard.ArchX64:   "https://code.visualstudio.com/sha/download?build=insider&os=darwin-universal",
						repo_standard.ArchARM64: "https://code.visualstudio.com/sha/download?build=insider&os=darwin-universal",
					},
				},
			},
		},
	}

}
