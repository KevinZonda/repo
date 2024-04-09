package entertainrepo

import "github.com/KevinZonda/repo/repo_standard"

const CHROME_MAC = "https://dl.google.com/chrome/mac/universal/stable/GGRO/googlechrome.dmg"
const CHROME_WIN = "https://dl.google.com/tag/s/appguid%3D%7B8A69D345-D564-463C-AFF1-A69D9E530F96%7D%26iid%3D%7B886FB653-BD28-B159-6AD2-4DD7724E9E75%7D%26lang%3Den%26browser%3D4%26usagestats%3D0%26appname%3DGoogle%2520Chrome%26needsadmin%3Dprefers%26ap%3Dx64-stable-statsdef_1%26installdataindex%3Dempty/update2/installers/ChromeSetup.exe"

type Chrome struct {
	repo_standard.BaseSync
}

func (t Chrome) GetPackage() repo_standard.Package {
	return repo_standard.Package{
		DisplayName: "Chrome",
		UName:       "chrome",
		Category:    "Entertainment",
		Versions: repo_standard.Latest(map[repo_standard.Platform]map[repo_standard.Arch]string{
			repo_standard.PlatformWin: {
				repo_standard.ArchX64: CHROME_WIN,
			},
			repo_standard.PlatformMac: {
				repo_standard.ArchX64:   CHROME_MAC,
				repo_standard.ArchARM64: CHROME_MAC,
			},
			//repo_standard.PlatformLinux: {},
		}),
	}.VersionToHistory()

}
