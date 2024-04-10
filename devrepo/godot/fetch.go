package godot

import "github.com/KevinZonda/repo/utils"

const GODOT_VERSION_URL = "https://raw.githubusercontent.com/godotengine/godot-website/master/_data/versions.yml"

func fetchIndex() (GDVersions, error) {

	vf, err := utils.HttpGetYaml[GDVersions](GODOT_VERSION_URL)
	return vf, err
}
