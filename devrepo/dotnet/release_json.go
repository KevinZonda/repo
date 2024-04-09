package dotnet

import "github.com/KevinZonda/repo/utils"

const RELEASES_INDEX = "https://dotnetcli.blob.core.windows.net/dotnet/release-metadata/releases-index.json"

type ReleaseIndexJson struct {
	Schema        string               `json:"$schema"`
	ReleasesIndex []ReleasesIndexBrief `json:"releases-index"`
}
type ReleasesIndexBrief struct {
	ChannelVersion    string           `json:"channel-version"`
	LatestRelease     string           `json:"latest-release"`
	LatestReleaseDate utils.Date       `json:"latest-release-date"`
	Security          bool             `json:"security"`
	LatestRuntime     string           `json:"latest-runtime"`
	LatestSdk         string           `json:"latest-sdk"`
	Product           string           `json:"product"`
	ReleaseType       ReleaseType      `json:"release-type"`
	SupportPhase      SupportPhaseType `json:"support-phase"`
	ReleasesJSON      string           `json:"releases.json"`
	EolDate           utils.Date       `json:"eol-date"`
}

type ReleaseIndexDetail struct {
	ChannelVersion    string           `json:"channel-version"`
	LatestRelease     string           `json:"latest-release"`
	LatestReleaseDate utils.Date       `json:"latest-release-date"`
	LatestRuntime     string           `json:"latest-runtime"`
	LatestSdk         string           `json:"latest-sdk"`
	SupportPhase      SupportPhaseType `json:"support-phase"`
	ReleaseType       ReleaseType      `json:"release-type"`
	LifecyclePolicy   string           `json:"lifecycle-policy"`
	Releases          []DotnetRelease  `json:"releases"`
}

type SupportPhaseType string

const (
	SupportPhaseTypeEOL         SupportPhaseType = "eof"
	SupportPhaseTypeActive      SupportPhaseType = "active"
	SupportPhaseTypeMaintenance SupportPhaseType = "maintenance"
	SupportPhaseTypePreview     SupportPhaseType = "preview"
)

type ReleaseType string

const (
	ReleaseTypeSts ReleaseType = "sts"
	ReleaseTypeLts ReleaseType = "lts"
)

type DotnetRelease struct {
	ReleaseDate       utils.Date            `json:"release-date"`
	ReleaseVersion    string                `json:"release-version"`
	Security          bool                  `json:"security"`
	CveList           []any                 `json:"cve-list"`
	ReleaseNotes      string                `json:"release-notes"`
	Runtime           RuntimeInfo           `json:"runtime"`
	Sdk               SdkInfo               `json:"sdk"`
	Sdks              []SdkInfo             `json:"sdks"`
	AspnetcoreRuntime AspnetCoreRelease     `json:"aspnetcore-runtime"`
	Windowsdesktop    WindowsDesktopRelease `json:"windowsdesktop"`
}

type DownloadFile struct {
	Name  string `json:"name"`
	Rid   string `json:"rid"`
	URL   string `json:"url"`
	Hash  string `json:"hash"`
	Akams string `json:"akams,omitempty"`
}

type SdkInfo struct {
	Version        string         `json:"version"`
	VersionDisplay string         `json:"version-display"`
	RuntimeVersion string         `json:"runtime-version"`
	VsVersion      string         `json:"vs-version"`
	VsMacVersion   string         `json:"vs-mac-version"`
	VsSupport      string         `json:"vs-support"`
	VsMacSupport   string         `json:"vs-mac-support"`
	CsharpVersion  string         `json:"csharp-version"`
	FsharpVersion  string         `json:"fsharp-version"`
	VbVersion      string         `json:"vb-version"`
	Files          []DownloadFile `json:"files"`
}
type RuntimeInfo struct {
	Version        string         `json:"version"`
	VersionDisplay string         `json:"version-display"`
	VsSupport      string         `json:"vs-support"`
	VsMacVersion   string         `json:"vs-mac-version"`
	Files          []DownloadFile `json:"files"`
}

type AspnetCoreRelease struct {
	Version                 string         `json:"version"`
	VersionDisplay          string         `json:"version-display"`
	VersionAspnetcoremodule []string       `json:"version-aspnetcoremodule"`
	VsVersion               string         `json:"vs-version"`
	Files                   []DownloadFile `json:"files"`
}

type WindowsDesktopRelease struct {
	Version        string         `json:"version"`
	VersionDisplay string         `json:"version-display"`
	Files          []DownloadFile `json:"files"`
}
