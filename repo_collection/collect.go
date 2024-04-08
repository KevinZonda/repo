package repo_collection

import (
	_go "github.com/KevinZonda/repo/devrepo/go"
	"github.com/KevinZonda/repo/devrepo/node"
	"github.com/KevinZonda/repo/entertainrepo"
	"github.com/KevinZonda/repo/repo_standard"
	"log"
	"sync"
	"time"
)

var R = Collection{
	entertainrepo.Steam{},
	entertainrepo.Telegram{},
	&node.NodeRepo{},
	&_go.GoRepo{},
}

type Collection []repo_standard.IRepo

func (c Collection) Sync() {
	for _, r := range c {
		log.Println("Syncing", r.GetPackage().UName)
		r.Sync()
	}
	UpdatedAt = time.Now()
	c.syncCache()
}

var cache *repo_standard.FullRepository

func (c Collection) FullRepository() repo_standard.FullRepository {
	if cache == nil {
		c.syncCache()
	}
	return *cache
}

var l sync.Mutex

func (c Collection) syncCache() {
	log.Println("Syncing repo cache")
	l.Lock()
	defer l.Unlock()
	m := c.fullRepo()
	cache = &m
}

func (c Collection) fullRepo() repo_standard.FullRepository {
	pkgs := make(map[string]repo_standard.Package)
	for _, r := range c {
		pkgs[r.GetPackage().UName] = r.GetPackage()
	}
	return repo_standard.FullRepository{
		Updated:  UpdatedAt,
		Packages: pkgs,
	}
}

var UpdatedAt time.Time

func (c Collection) StartDaemon(dur time.Duration) <-chan bool {
	stopCh := make(chan bool)
	go c.Sync()
	go func() {
		for {
			select {
			case <-time.After(dur):
				c.Sync()
			case <-stopCh:
				close(stopCh)
			}
		}
	}()
	return stopCh

}
