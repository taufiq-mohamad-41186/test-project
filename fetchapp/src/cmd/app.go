package main

import (
	"github.com/allegro/bigcache"
	"github.com/gorilla/mux"
	domain "github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/domain"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/usecase"
	restHandler "github.com/taufiq-mohamad-41186/test-project/fetchapp/src/handler/rest"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
	"time"
)

var (
	conf Conf

	httpClient http.Client

	// Business Layer
	dom *domain.Domain
	uc  *usecase.Usecase
)

type Conf struct {
	Auth restHandler.AuthOptions
	Business
}
type Business struct {
	Domain  domain.Options
	Usecase usecase.Options
}

func Serve() {
	b, err := os.ReadFile("etc/conf/conf.yaml")
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(b, &conf); err != nil {
		panic(err)
	}

	httpClient = http.Client{}
	bCache, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             1024,
		LifeWindow:         1 * time.Hour,
		CleanWindow:        5 * time.Minute,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       500,
		Verbose:            false,
		HardMaxCacheSize:   256,
		OnRemove:           nil,
		OnRemoveWithReason: nil,
	})
	if err != nil {
		panic(err)
	}

	dom = domain.Init(httpClient, bCache, conf.Business.Domain)
	uc = usecase.Init(conf.Business.Usecase, dom)

	auth := restHandler.InitAuth(conf.Auth)
	_ = restHandler.Init(auth, mux.NewRouter(), uc)
}
