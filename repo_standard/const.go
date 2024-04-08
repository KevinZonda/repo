package repo_standard

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

type IRepo interface {
	ISync
	GetPackage() Package
}

type ISync interface {
	Sync()
}

type BaseSync struct {
}

func (b BaseSync) Sync() {}
