package git

import (
	"github.com/KevinZonda/repo/utils"
	"regexp"
	"strings"
)

const URL = "https://git-scm.com/download/win"

type index struct {
	winx64  string
	winx86  string
	version string
}

var r, _ = regexp.Compile(`https://github\.com/git-for-windows/git/releases/download/.*/Git.*\.exe`)

func fetchIndex() (index, error) {
	body, err := utils.HttpGetStr(URL)
	if err != nil {
		return index{}, err
	}
	matches := flatten(r.FindAllStringSubmatch(body, -1))
	x86 := ""
	x64 := ""
	for _, v := range matches {
		if strings.Contains(v, "32-bit") {
			if x86 == "" {
				x86 = v
			}
			continue
		}
		if strings.Contains(v, "64-bit") {
			if x64 == "" {
				x64 = v
			}
			continue
		}
	}
	vstr := utils.NotEmptyOne(x86, x64)
	if vstr != "" {
		urls := strings.Split(vstr, "/")
		v := urls[len(urls)-1]
		urls = strings.Split(v, "-")
		if len(urls) < 3 {
			vstr = ""
		} else {
			vstr = urls[1]
		}
	}
	return index{x64, x86, vstr}, nil
}

func flatten(s [][]string) []string {
	var r []string
	for _, v := range s {
		r = append(r, v...)
	}
	return r
}
