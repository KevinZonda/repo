package _go

import (
	"github.com/KevinZonda/repo/repo_standard"
	"github.com/KevinZonda/repo/utils"
	"strings"
)

const GO_INDEX_URL = "https://go.dev/dl/?mode=json&include=all"

func fetchIndex() (GoIndex, error) {
	index, err := utils.HttpGetJson[GoIndex](GO_INDEX_URL)
	return index, err
}

type GoIndex []GoIndexItem

func (i GoIndex) Stable() GoIndexItem {
	for _, item := range i {
		if item.Stable {
			return item
		}
	}
	return GoIndexItem{}
}

func (i GoIndex) Latest() GoIndexItem {
	return i[0]
}

type GoIndexItem struct {
	Version string        `json:"version"`
	Stable  bool          `json:"stable"`
	Files   []GoIndexFile `json:"files"`
}

func (i GoIndexItem) FileList() GoIndexFileMap {
	m := map[string]GoIndexFile{}
	for _, f := range i.Files {
		m[f.Filename] = f
	}
	return m
}

type GoIndexFileMap map[string]GoIndexFile

type GoIndexFile struct {
	Filename string `json:"filename"`
	Os       string `json:"os"`
	Arch     string `json:"arch"`
	Version  string `json:"version"`
	Sha256   string `json:"sha256"`
	Size     int    `json:"size"`
	Kind     string `json:"kind"`
}

func (m GoIndexFileMap) Url(filename string) string {
	_, ok := m[filename]
	if !ok {
		return ""
	}
	return "https://go.dev/dl/" + filename
}

func (i GoIndexItem) VersionedUrl() repo_standard.VersionedUrl {
	fl := i.FileList()
	m := map[repo_standard.Platform]map[repo_standard.Arch]string{
		repo_standard.PlatformWin: {
			repo_standard.ArchX64:   fl.Url(i.versionedFilename("windows-amd64")),
			repo_standard.ArchX86:   fl.Url(i.versionedFilename("windows-386")),
			repo_standard.ArchARM64: fl.Url(i.versionedFilename("windows-arm64")),
		},
		repo_standard.PlatformMac: {
			repo_standard.ArchX64:   fl.Url(i.versionedFilename("darwin-amd64")),
			repo_standard.ArchARM64: fl.Url(i.versionedFilename("darwin-arm64")),
		},
		repo_standard.PlatformLinux: {
			repo_standard.ArchX64:   fl.Url(i.versionedFilename("linux-amd64")),
			repo_standard.ArchX86:   fl.Url(i.versionedFilename("linux-386")),
			repo_standard.ArchARM64: fl.Url(i.versionedFilename("linux-amd64")),
		},
	}
	return repo_standard.VersionedUrl{
		Version: i.Version,
		Urls:    m,
	}
}

func (i GoIndexItem) versionedFilename(arch string) string {
	ext := ".tar.gz"
	if strings.HasPrefix(arch, "win") {
		ext = ".msi"
	} else if strings.HasPrefix(arch, "darwin") {
		ext = ".pkg"
	}
	return i.Version + "." + arch + ext
}
