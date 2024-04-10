package godot

import (
	"github.com/KevinZonda/repo/repo_standard"
	"log"
)

type GodotRepo struct {
	idx GDVersions
}

func (g *GodotRepo) GetPackage() repo_standard.Package {
	pkg := repo_standard.Package{
		DisplayName: "Godot",
		UName:       "godot",
		Category:    "Development",
	}
	if len(g.idx) == 0 {
		return pkg
	}
	pkg.History, pkg.Versions = g.idx.ToVersions()
	return pkg
}

func (g *GodotRepo) Sync() {
	idx, err := fetchIndex()
	if err != nil {
		log.Println(err)
		log.Println("Failed to sync Node repo")
		return
	}
	g.idx = idx
}
