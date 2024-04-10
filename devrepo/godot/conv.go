package godot

import "strings"

// https://github.com/godotengine/godot-builds/releases/download/4.2.1-stable/Godot_v4.2.1-stable_win64.exe.zip
// https://github.com/godotengine/godot-builds/releases/download/4.2.1-stable/Godot_v4.2.1-stable_mono_win64.exe.zip
// https://github.com/godotengine/godot-builds/releases/download/4.2.1-stable/Godot_v4.2.1-stable_mono_macos.universal.zip

// https://github.com/godotengine/godot-builds/releases/download/4.2.1-stable/Godot_v4.2.1-stable_mono_linux.x86_64.zip

type GDPlatform string

const (
	GD_PLATFORM_WIN   GDPlatform = "win"
	GD_PLATFORM_LINUX GDPlatform = "linux"
	GD_PLATFORM_MAC   GDPlatform = "macos"
)

type GDArch string

const (
	GD_ARCH_LINUX_X86_64 GDArch = "_x86_64"
	GD_ARCH_LINUX_ARM64  GDArch = "_arm64"
	GD_ARCH_LINUX_X86_32 GDArch = "_x86_32"
	GD_ARCH_MAC_UNIV     GDArch = ".universal"
	GD_ARCH_WIN_32       GDArch = "32"
	GD_ARCH_WIN_64       GDArch = "64"
)

func filename(version string, flavour string, platform GDPlatform, arch GDArch, isMono bool) string {
	// else
	//      # Format was slightly different up until 2.1.
	//      if version_bits[0] == "1" or (version_bits[0] == "2" and version_bits[1] == "0")
	//        download_file = "Godot_v#{version_name}_#{version_flavor}_#{download_slug}"
	//                                              ^^^
	//      else
	//        download_file = "Godot_v#{version_name}-#{version_flavor}_#{download_slug}"
	//      end
	//    end
	download_slug := string(platform) + string(arch)
	mono := "_"
	if isMono {
		mono = "_mono_"
	}
	sep := "-"
	if strings.HasPrefix(version, "0.") || strings.HasPrefix(version, "1.") || strings.HasPrefix(version, "2.") {
		sep = "_"
	}
	return "Godot_v" + version + sep + flavour + mono + download_slug + ".zip"
}

func isSupportMono(version string) bool {
	return !strings.HasPrefix(version, "0.") && !strings.HasPrefix(version, "1.") && !strings.HasPrefix(version, "2.")
}
