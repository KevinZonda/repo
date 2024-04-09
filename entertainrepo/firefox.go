package entertainrepo

import "github.com/KevinZonda/repo/repo_standard"

const FIREFOX_MAC = "https://download.mozilla.org/?product=firefox-latest-ssl&os=osx&lang=en-GB"
const FIREFOX_WIN = "https://download.mozilla.org/?product=firefox-latest-ssl&os=win&lang=en-GB"

type Firefox struct {
	repo_standard.BaseSync
}

func (t Firefox) GetPackage() repo_standard.Package {
	stable := repo_standard.VersionedUrl{
		Version: "stable",
		Urls: map[repo_standard.Platform]map[repo_standard.Arch]string{
			repo_standard.PlatformWin: {
				repo_standard.ArchX64: FIREFOX_WIN,
			},
			repo_standard.PlatformMac: {
				repo_standard.ArchX64:   FIREFOX_MAC,
				repo_standard.ArchARM64: FIREFOX_MAC,
			},
		},
	}

	return repo_standard.Package{
		DisplayName: "Firefox",
		UName:       "firefox",
		Category:    "Entertainment",
		Versions: repo_standard.VersionedUrls{
			"stable": stable,
			"latest": stable,
		},
	}.VersionToHistory()

}
