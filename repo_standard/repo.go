package repo_standard

type VersionedUrls map[string]VersionedUrl

type VersionedUrl struct {
	Version string
	Urls    map[Platform]map[Arch]string
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
	DisplayName string
	UName       string
	Category    string
	Urls        VersionedUrls
}

type FullRepository struct {
	Packages map[string]Package
}
