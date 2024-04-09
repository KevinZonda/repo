package entertainrepo

import "github.com/KevinZonda/repo/repo_standard"

const CHROME_MAC = "https://dl.google.com/chrome/mac/universal/stable/GGRO/googlechrome.dmg"
const CHROME_WIN = "https://dl.google.com/tag/s/appguid%3D%7B8A69D345-D564-463C-AFF1-A69D9E530F96%7D%26iid%3D%7B886FB653-BD28-B159-6AD2-4DD7724E9E75%7D%26lang%3Den%26browser%3D4%26usagestats%3D0%26appname%3DGoogle%2520Chrome%26needsadmin%3Dprefers%26ap%3Dx64-stable-statsdef_1%26installdataindex%3Dempty/update2/installers/ChromeSetup.exe"

const CHROME_CANARY_WIN = "https://dl.google.com/tag/s/appguid%3D%7B4EA16AC7-FD5A-47C3-875B-DBF4A2008C20%7D%26iid%3D%7B68DE4A80-1350-9DCA-9861-DD9568B38B9C%7D%26lang%3Den-GB%26browser%3D4%26usagestats%3D0%26appname%3DGoogle%2520Chrome%2520Canary%26needsadmin%3Dfalse%26ap%3Dx64-canary-statsdef_1%26installdataindex%3Dempty/update2/installers/ChromeSetup.exe"
const CHROME_CANARY_MAC = "https://dl.google.com/chrome/mac/universal/canary/googlechromecanary.dmg"

const CHROME_BETA_WIN = "https://dl.google.com/tag/s/appguid%3D%7B8237E44A-0054-442C-B6B6-EA0509993955%7D%26iid%3D%7B68DE4A80-1350-9DCA-9861-DD9568B38B9C%7D%26lang%3Den-GB%26browser%3D4%26usagestats%3D0%26appname%3DGoogle%2520Chrome%2520Beta%26needsadmin%3Dprefers%26ap%3D-arch_x64-statsdef_1%26installdataindex%3Dempty/update2/installers/ChromeSetup.exe"
const CHROME_BETA_MAC = "https://dl.google.com/chrome/mac/universal/beta/googlechromebeta.dmg"

const CHROME_DEV_WIN = "https://dl.google.com/tag/s/appguid%3D%7B401C381F-E0DE-4B85-8BD8-3F3F14FBDA57%7D%26iid%3D%7B68DE4A80-1350-9DCA-9861-DD9568B38B9C%7D%26lang%3Den-GB%26browser%3D4%26usagestats%3D0%26appname%3DGoogle%2520Chrome%2520Dev%26needsadmin%3Dprefers%26ap%3D-arch_x64-statsdef_1%26installdataindex%3Dempty/update2/installers/ChromeSetup.exe"
const CHROME_DEV_MAC = "https://dl.google.com/chrome/mac/universal/dev/googlechromedev.dmg"

type Chrome struct {
	repo_standard.BaseSync
}

func (t Chrome) GetPackage() repo_standard.Package {
	stable := repo_standard.VersionedUrl{
		Version: "stable",
		Urls: map[repo_standard.Platform]map[repo_standard.Arch]string{
			repo_standard.PlatformWin: {
				repo_standard.ArchX64: CHROME_WIN,
			},
			repo_standard.PlatformMac: {
				repo_standard.ArchX64:   CHROME_MAC,
				repo_standard.ArchARM64: CHROME_MAC,
			},
		},
	}

	canary := repo_standard.VersionedUrl{
		Version: "canary",
		Urls: map[repo_standard.Platform]map[repo_standard.Arch]string{
			repo_standard.PlatformWin: {
				repo_standard.ArchX64: CHROME_CANARY_WIN,
			},
			repo_standard.PlatformMac: {
				repo_standard.ArchX64:   CHROME_CANARY_MAC,
				repo_standard.ArchARM64: CHROME_CANARY_MAC,
			},
		},
	}

	beta := repo_standard.VersionedUrl{
		Version: "beta",
		Urls: map[repo_standard.Platform]map[repo_standard.Arch]string{
			repo_standard.PlatformWin: {
				repo_standard.ArchX64: CHROME_BETA_WIN,
			},
			repo_standard.PlatformMac: {
				repo_standard.ArchX64:   CHROME_BETA_MAC,
				repo_standard.ArchARM64: CHROME_BETA_MAC,
			},
		},
	}

	dev := repo_standard.VersionedUrl{
		Version: "dev",
		Urls: map[repo_standard.Platform]map[repo_standard.Arch]string{
			repo_standard.PlatformWin: {
				repo_standard.ArchX64: CHROME_DEV_WIN,
			},
			repo_standard.PlatformMac: {
				repo_standard.ArchX64:   CHROME_DEV_MAC,
				repo_standard.ArchARM64: CHROME_DEV_MAC,
			},
		},
	}

	return repo_standard.Package{
		DisplayName: "Chrome",
		UName:       "chrome",
		Category:    "Entertainment",
		Versions: repo_standard.VersionedUrls{
			"stable": stable,
			"canary": canary,
			"beta":   beta,
			"dev":    dev,
			"latest": stable,
		},
	}.VersionToHistory()

}
