package entertainrepo

import "github.com/KevinZonda/repo/repo_standard"

var SteamDl = map[string]string{
	"win":    "https://cdn.cloudflare.steamstatic.com/client/installer/SteamSetup.exe",
	"mac":    "https://cdn.cloudflare.steamstatic.com/client/installer/steam.dmg",
	"debian": "https://cdn.cloudflare.steamstatic.com/client/installer/steam.dev",
}

type Steam struct {
	repo_standard.BaseSync
}

func (s Steam) GetPackage() repo_standard.Package {
	return repo_standard.Package{
		DisplayName: "Steam",
		UName:       "steam",
		Category:    "Entertainment",
		Versions: repo_standard.Latest(map[repo_standard.Platform]map[repo_standard.Arch]string{
			repo_standard.PlatformWin: {
				repo_standard.ArchX64:   SteamDl["win"],
				repo_standard.ArchARM64: SteamDl["win"],
				repo_standard.ArchX86:   SteamDl["win"],
			},
			repo_standard.PlatformMac: {
				repo_standard.ArchX64:   SteamDl["mac"],
				repo_standard.ArchARM64: SteamDl["mac"],
			},
		}),
	}

}
