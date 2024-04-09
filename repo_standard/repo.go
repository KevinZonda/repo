package repo_standard

import "time"

type VersionedUrls map[string]VersionedUrl

type VersionedUrl struct {
	Version string                       `json:"version"`
	Urls    map[Platform]map[Arch]string `json:"urls"`
}

func Latest(versionedUrls map[Platform]map[Arch]string) VersionedUrls {
	return map[string]VersionedUrl{
		"latest": {
			Version: "latest",
			Urls:    versionedUrls,
		},
	}
}

func NewUrls() VersionedUrl {
	return VersionedUrl{Urls: map[Platform]map[Arch]string{}}
}

type ComposedName string

//func ComposeName(platform Platform, arch Arch) ComposedName {
//	return ComposedName(string(platform) + "-" + string(arch))
//}
//
//func (c ComposedName) Split() (Platform, Arch) {
//	v := strings.Split(string(c), "-")
//	return Platform(v[0]), Arch(v[1])
//}
//
//func (c ComposedName) Arch() Arch {
//	_, arch := c.Split()
//	return arch
//}
//
//func (c ComposedName) Platform() Platform {
//	platform, _ := c.Split()
//	return platform
//}
//
//func (c *ComposedName) UnmarshalJSON(bytes []byte) error {
//	*c = ComposedName(strings.Trim(string(bytes), `"`))
//	return nil
//}

type Package struct {
	DisplayName string        `json:"display_name"`
	UName       string        `json:"uid"`
	Category    string        `json:"category"`
	Versions    VersionedUrls `json:"versions"`
	History     VersionedUrls `json:"history,omitempty"`
}

func (p Package) WithoutHistory() Package {
	p.History = nil
	return p
}

func (p Package) VersionToHistory() Package {
	p.History = p.Versions
	return p
}

type FullRepository struct {
	Updated  time.Time          `json:"updated"`
	Packages map[string]Package `json:"packages"`
}
