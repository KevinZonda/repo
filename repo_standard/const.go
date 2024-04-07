package repo_standard

import "strings"

type Platform string

const (
	PlatformWin   Platform = "win"
	PlatformMac   Platform = "mac"
	PlatformLinux Platform = "linux"
)

type Arch string

const (
	ArchX64   Arch = "x64"
	ArchX86   Arch = "x86"
	ArchARM64 Arch = "arm64"
	ArchARMv7 Arch = "armv7"
)
