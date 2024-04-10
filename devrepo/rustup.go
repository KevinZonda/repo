package devrepo

import "github.com/KevinZonda/repo/repo_standard"

type RustupRepo struct {
	repo_standard.BaseSync
}

func (v RustupRepo) GetPackage() repo_standard.Package {
	m := map[repo_standard.Platform]map[repo_standard.Arch]string{
		repo_standard.PlatformWin: {
			repo_standard.ArchX64: "https://static.rust-lang.org/rustup/dist/i686-pc-windows-gnu/rustup-init.exe",
		},
		repo_standard.PlatformMac: {
			repo_standard.ArchX64:   "https://static.rust-lang.org/rustup/rustup-init.sh",
			repo_standard.ArchARM64: "https://static.rust-lang.org/rustup/rustup-init.sh",
		},
		repo_standard.PlatformLinux: {
			repo_standard.ArchX64:   "https://static.rust-lang.org/rustup/rustup-init.sh",
			repo_standard.ArchARM64: "https://static.rust-lang.org/rustup/rustup-init.sh",
		},
	}
	return repo_standard.Package{
		DisplayName: "Rustup",
		UName:       "rustup",
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
		},
	}.VersionToHistory()
}
