package entertainrepo

import "github.com/KevinZonda/repo/repo_standard"

var TelegramDl = map[string]string{
	"win64":          "https://telegram.org/dl/desktop/win64",
	"win64_portable": "https://telegram.org/dl/desktop/win64_portable",
	"win":            "https://telegram.org/dl/desktop/win",
	"win_portable":   "https://telegram.org/dl/desktop/win_portable",
	"mac":            "https://telegram.org/dl/desktop/mac",
	"mac_store":      "https://telegram.org/dl/desktop/mac_store",
	"linux":          "https://telegram.org/dl/desktop/linux",
	"linux_flatpak":  "https://telegram.org/dl/desktop/flatpak",
	"linux_snap":     "https://telegram.org/dl/desktop/snap",
}

type Telegram struct {
	repo_standard.BaseSync
}

func (t Telegram) GetPackage() repo_standard.Package {
	return repo_standard.Package{
		DisplayName: "Telegram",
		UName:       "telegram",
		Category:    "Entertainment",
		Versions: repo_standard.Latest(map[repo_standard.Platform]map[repo_standard.Arch]string{
			repo_standard.PlatformWin: {
				repo_standard.ArchX64: TelegramDl["win64"],
				repo_standard.ArchX86: TelegramDl["win"],
			},
			repo_standard.PlatformMac: {
				repo_standard.ArchX64:   TelegramDl["mac"],
				repo_standard.ArchARM64: TelegramDl["mac"],
			},
			repo_standard.PlatformLinux: {
				repo_standard.ArchX64: TelegramDl["linux"],
			},
		}),
	}

}
