package godot

import "github.com/KevinZonda/repo/repo_standard"

type GDVersions []struct {
	Name         string      `yaml:"name"`
	Flavor       string      `yaml:"flavor"`
	ReleaseDate  string      `yaml:"release_date"`
	ReleaseNotes string      `yaml:"release_notes"`
	Releases     []GDRelease `yaml:"releases,omitempty"`
	Featured     string      `yaml:"featured,omitempty"`
}

type GDRelease struct {
	Name         string `yaml:"name"`
	ReleaseDate  string `yaml:"release_date"`
	ReleaseNotes string `yaml:"release_notes"`
}

func (vs GDVersions) ToVersions() (all repo_standard.VersionedUrls, simple repo_standard.VersionedUrls) {
	u := make(repo_standard.VersionedUrls)
	simple = make(repo_standard.VersionedUrls)
	// - name: "4.2"
	//  flavor: "stable"
	//  releases:
	//    - name: "rc2"
	//    - name: "rc1"
	//    - name: "beta6"
	// first one is
	stable := false // set stable
	latest := false
	for _, v := range vs {
		version := v.Name
		if !latest {
			latest = true
			latestV := m(version, v.Flavor, false)
			latestM := m(version, v.Flavor, true)
			u["latest"], u["latest-mono"] = latestV, latestM
			simple["latest"], simple["latest-mono"] = latestV, latestM
		}
		if v.Flavor == "stable" && !stable {
			stable = true
			latestV := m(version, "stable", false)
			latestM := m(version, "stable", true)
			u["stable"], u["stable-mono"] = latestV, latestM
			simple["stable"], simple["stable-mono"] = latestV, latestM
		}

		root := urlmap(version, v.Flavor)
		for key, val := range root {
			u[key] = val
		}
		for _, release := range v.Releases {
			releaseM := urlmap(version, release.Name)
			for key, val := range releaseM {
				u[key] = val
			}
		}
	}
	return u, simple
}

func urlmap(version, flavour string) repo_standard.VersionedUrls {
	x := make(repo_standard.VersionedUrls)
	vanilla := m(version, flavour, false)
	x[version+"-"+flavour] = vanilla
	if flavour == "stable" {
		x[version] = vanilla
	}
	if isSupportMono(version) {
		mono := m(version, flavour, true)
		x[version+"-"+flavour+"-mono"] = mono
		if flavour == "stable" {
			x[version+"-mono"] = mono
		}
	}
	return x
}

func m(version, flavour string, mono bool) repo_standard.VersionedUrl {
	const DL_URL = "https://github.com/godotengine/godot-builds/releases/download/"
	// 	// https://github.com/godotengine/godot-builds/releases/download/4.2.1-stable/Godot_v4.2.1-stable_mono_linux_arm64.zip
	param := version + "-" + flavour + "/"
	URL := DL_URL + param
	return repo_standard.VersionedUrl{
		Version: version + "-" + flavour,
		Urls: map[repo_standard.Platform]map[repo_standard.Arch]string{
			repo_standard.PlatformWin: {
				repo_standard.ArchX64: URL + filename(version, flavour, GD_PLATFORM_WIN, GD_ARCH_WIN_64, mono),
				repo_standard.ArchX86: URL + filename(version, flavour, GD_PLATFORM_WIN, GD_ARCH_WIN_32, mono),
			},
			repo_standard.PlatformMac: {
				repo_standard.ArchX64:   URL + filename(version, flavour, GD_PLATFORM_MAC, GD_ARCH_MAC_UNIV, mono),
				repo_standard.ArchARM64: URL + filename(version, flavour, GD_PLATFORM_MAC, GD_ARCH_MAC_UNIV, mono),
			},
			repo_standard.PlatformLinux: {
				repo_standard.ArchX64:   URL + filename(version, flavour, GD_PLATFORM_LINUX, GD_ARCH_LINUX_X86_64, mono),
				repo_standard.ArchX86:   URL + filename(version, flavour, GD_PLATFORM_LINUX, GD_ARCH_LINUX_X86_32, mono),
				repo_standard.ArchARM64: URL + filename(version, flavour, GD_PLATFORM_LINUX, GD_ARCH_LINUX_ARM64, mono),
			},
		},
	}
}
