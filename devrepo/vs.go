package devrepo

import "github.com/KevinZonda/repo/repo_standard"

const VS_PRO = "https://c2rsetup.officeapps.live.com/c2r/downloadVS.aspx?sku=professional&channel=Release&version=VS2022"
const VS_COM = "https://c2rsetup.officeapps.live.com/c2r/downloadVS.aspx?sku=community&channel=Release&version=VS2022"
const VS_ENT = "https://c2rsetup.officeapps.live.com/c2r/downloadVS.aspx?sku=enterprise&channel=Release&version=VS2022"

type VSRepo struct {
	repo_standard.BaseSync
}

func (v VSRepo) vurl(sku string, url string) repo_standard.VersionedUrl {
	return repo_standard.VersionedUrl{
		Version: sku,
		Urls: map[repo_standard.Platform]map[repo_standard.Arch]string{
			repo_standard.PlatformWin: {
				repo_standard.ArchX64: url,
			},
		},
	}
}

func (v VSRepo) GetPackage() repo_standard.Package {
	return repo_standard.Package{
		DisplayName: "Visual Studio 2022 (VS)",
		UName:       "vs",
		Category:    "Development",
		Versions: map[string]repo_standard.VersionedUrl{
			"stable": v.vurl("community", VS_COM),
			"latest": v.vurl("community", VS_COM),

			"pro": v.vurl("pro", VS_PRO),
			"com": v.vurl("com", VS_COM),
			"ent": v.vurl("ent", VS_ENT),

			"community":    v.vurl("com", VS_COM),
			"professional": v.vurl("ins", VS_PRO),
			"enterprise":   v.vurl("ins", VS_ENT),
		},
		History: map[string]repo_standard.VersionedUrl{
			"stable": v.vurl("community", VS_COM),

			"pro":          v.vurl("pro", VS_PRO),
			"com":          v.vurl("com", VS_COM),
			"ent":          v.vurl("ent", VS_ENT),
			"community":    v.vurl("com", VS_COM),
			"professional": v.vurl("ins", VS_PRO),
			"enterprise":   v.vurl("ins", VS_ENT),
		},
	}
}
