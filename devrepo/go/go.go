package _go

import (
	"encoding/json"
	"github.com/KevinZonda/repo/repo_standard"
	"net/http"
	"strings"
)

const GO_INDEX_URL = "https://go.dev/dl/?mode=json&include=all"

type Repo struct {
	hc *http.Client
}

func NewRepo() *Repo {
	return &Repo{hc: http.DefaultClient}
}

func (r *Repo) fetchIndex() (GoIndex, error) {
	resp, err := r.hc.Get(GO_INDEX_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var index GoIndex
	err = json.NewDecoder(resp.Body).Decode(&index)
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

type GoIndexFile struct {
	Filename string `json:"filename"`
	Os       string `json:"os"`
	Arch     string `json:"arch"`
	Version  string `json:"version"`
	Sha256   string `json:"sha256"`
	Size     int    `json:"size"`
	Kind     string `json:"kind"`
}

func (i GoIndexItem) VersionedUrl() repo_standard.VersionedUrl {
	m := map[repo_standard.Platform]map[repo_standard.Arch]string{
		repo_standard.PlatformWin: {
			repo_standard.ArchX64:   dlUrl(i.versionedFilename("windows-amd64")),
			repo_standard.ArchX86:   dlUrl(i.versionedFilename("windows-386")),
			repo_standard.ArchARM64: dlUrl(i.versionedFilename("windows-arm64")),
		},
		repo_standard.PlatformMac: {
			repo_standard.ArchX64:   dlUrl(i.versionedFilename("darwin-amd64")),
			repo_standard.ArchARM64: dlUrl(i.versionedFilename("darwin-arm64")),
		},
		repo_standard.PlatformLinux: {
			repo_standard.ArchX64:   dlUrl(i.versionedFilename("linux-amd64")),
			repo_standard.ArchX86:   dlUrl(i.versionedFilename("linux-386")),
			repo_standard.ArchARM64: dlUrl(i.versionedFilename("linux-amd64")),
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

func dlUrl(filename string) string {
	return "https://go.dev/dl/" + filename
}
