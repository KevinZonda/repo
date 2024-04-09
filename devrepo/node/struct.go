package node

import (
	"github.com/KevinZonda/repo/repo_standard"
	"github.com/KevinZonda/repo/utils"
)

type IndexJson []IndexItem

type IndexItem struct {
	Version  string     `json:"version"`
	Date     utils.Date `json:"date"`
	Files    []string   `json:"files"`
	Npm      string     `json:"npm,omitempty"`
	V8       string     `json:"v8"`
	Uv       string     `json:"uv,omitempty"`
	Zlib     string     `json:"zlib,omitempty"`
	Openssl  string     `json:"openssl,omitempty"`
	Modules  string     `json:"modules,omitempty"`
	Lts      any        `json:"lts"`
	Security bool       `json:"security"`
}

func (i IndexItem) IsLts() bool {
	switch lT := i.Lts.(type) {
	case bool:
		return lT
	case string:
		return true
	default:
		return false
	}
}

func (j IndexJson) Latest() IndexItem {
	if len(j) == 0 {
		return IndexItem{}
	}
	return j[0]
}

func (j IndexJson) LatestLts() IndexItem {
	for _, item := range j {
		if item.IsLts() {
			return item
		}
	}
	return IndexItem{}
}

func (i IndexItem) DownloadList() map[string]string {
	m := map[string]string{}
	for _, file := range i.Files {
		m[file] = "https://nodejs.org/dist/" + i.Version + "/" + filename(i.Version, file)
		// https://nodejs.org/dist/v16.13.0/node-v16.13.0-linux-x64.tar.xz
	}
	return m
}

func filename(version string, platform string) string {
	r, ok := rules[platform]
	if !ok {
		return ""
	}
	name := r.FileName
	if name != "" {
		name = "-" + name
	}
	return "node-" + version + name + r.Ext
}

type ConversionRule struct {
	FileName string
	Arch     string
	Ext      string
}

var rules = map[string]ConversionRule{
	"aix-ppc64":     {FileName: "aix-ppc64", Arch: "ppc64", Ext: ".tar.gz"},
	"headers":       {FileName: "headers", Arch: "", Ext: ".tar.gz"},
	"linux-arm64":   {FileName: "linux-arm64", Arch: "arm64", Ext: ".tar.gz"},
	"linux-armv7l":  {FileName: "linux-armv7l", Arch: "armv7l", Ext: ".tar.gz"},
	"linux-ppc64le": {FileName: "linux-ppc64le", Arch: "ppc64", Ext: ".tar.gz"},
	"linux-s390x":   {FileName: "linux-s390x", Arch: "s390x", Ext: ".tar.gz"},
	"linux-x64":     {FileName: "linux-x64", Arch: "x64", Ext: ".tar.gz"},
	"osx-arm64-tar": {FileName: "darwin-arm64", Arch: "arm64", Ext: ".tar.gz"},
	"osx-x64-pkg":   {FileName: "", Arch: "x64", Ext: ".pkg"},
	"osx-x64-tar":   {FileName: "darwin-x64", Arch: "x64", Ext: ".tar.gz"},
	"src":           {FileName: "", Arch: "", Ext: ".tar.gz"},
	"win-arm64-7z":  {FileName: "win-arm64", Arch: "arm64", Ext: ".7z"},
	"win-arm64-zip": {FileName: "win-arm64", Arch: "arm64", Ext: ".zip"},
	"win-arm64-msi": {FileName: "arm64", Arch: "arm64", Ext: ".msi"},
	"win-x64-7z":    {FileName: "win-x64", Arch: "x64", Ext: ".7z"},
	// "win-x64-exe":   {FileName: "win-x64", Arch: "x64", Ext: ".exe"},
	"win-x64-msi": {FileName: "x64", Arch: "x64", Ext: ".msi"},
	"win-x64-zip": {FileName: "win-x64", Arch: "x64", Ext: ".zip"},
	"win-x86-7z":  {FileName: "win-x86", Arch: "x86", Ext: ".7z"},
	// "win-x86-exe":   {FileName: "x86", Arch: "x86", Ext: ".exe"},
	"win-x86-msi": {FileName: "x86", Arch: "x86", Ext: ".msi"},
	"win-x86-zip": {FileName: "win-x86", Arch: "x86", Ext: ".zip"},
}

func (i IndexItem) ToVersionedUrl() repo_standard.VersionedUrl {
	dl_list := i.DownloadList()
	m := map[repo_standard.Platform]map[repo_standard.Arch]string{
		repo_standard.PlatformWin: {
			repo_standard.ArchX64:   tryGetWinUrl(dl_list, "win-x64"),
			repo_standard.ArchARM64: tryGetWinUrl(dl_list, "win-arm64"),
			repo_standard.ArchX86:   tryGetWinUrl(dl_list, "win-x86"),
		},
		repo_standard.PlatformMac: {
			repo_standard.ArchX64:   dl_list["osx-x64-pkg"],
			repo_standard.ArchARM64: dl_list["osx-arm64-tar"],
		},
		repo_standard.PlatformLinux: {
			repo_standard.ArchX64:   dl_list["linux-x64"],
			repo_standard.ArchARM64: dl_list["linux-arm64"],
		},
	}
	return repo_standard.VersionedUrl{
		Version: i.Version,
		Urls:    m,
	}
}

func tryGetWinUrl(dlList map[string]string, winPrefix string) string {
	seq := []string{"msi", "exe", "zip"}
	for _, v := range seq {
		if url, ok := dlList[winPrefix+"-"+v]; ok {
			return url
		}
	}
	return ""
}
