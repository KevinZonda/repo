package dotnet

import (
	"github.com/KevinZonda/repo/utils"
	"log"
)

func (f ReleaseIndexJson) fetch() map[string]ReleaseIndexDetail {
	m := map[string]ReleaseIndexDetail{}
	for _, v := range f.ReleasesIndex {
		i, e := utils.HttpGetJson[ReleaseIndexDetail](v.ReleasesJSON)
		if e != nil {
			log.Println(e)
		}
		m[i.ChannelVersion] = i
	}
	return m
}
