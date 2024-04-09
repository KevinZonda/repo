package msys2

import (
	"github.com/KevinZonda/repo/utils"
	"regexp"
)

const URL = "https://www.msys2.org/"

type index struct {
	winx64  string
	version string
}

var r, _ = regexp.Compile(`https://github\.com/msys2/msys2-installer/releases/download/[_\-0-9]+/msys2-x86_64[_\-0-9]+\.exe`)
var versionReg, _ = regexp.Compile(`(\d+)(?:\.exe)`)

func fetchIndex() (index, error) {
	body, err := utils.HttpGetStr(URL)
	if err != nil {
		return index{}, err
	}
	matches := flatten(r.FindAllStringSubmatch(body, -1))
	x64 := ""
	if len(matches) > 0 {
		x64 = matches[0]
	}

	//fmt.Println(matches)
	matches = versionReg.FindStringSubmatch(x64)
	version := ""
	switch len(matches) {
	case 1:
		version = matches[0]
	default:
		version = matches[1]
	}
	return index{x64, version}, nil
}

func flatten(s [][]string) []string {
	var r []string
	for _, v := range s {
		r = append(r, v...)
	}
	return r
}
