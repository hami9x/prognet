package main

import (
	"path"

	ng "github.com/gopherjs/go-angularjs"
	"github.com/gopherjs/gopherjs/js"
	"github.com/phaikawl/prognet/biz"
)

const (
	ServerAddress = ""
)

var (
	gLocalStorage *LocalStorage = InitLocalStorage()
	gSiteData                   = &siteData{}
)

type userData struct {
	Username string
	Token    string
}

type siteData struct {
	user *userData
}

func (d *siteData) User() *userData {
	if d.user != nil {
		return d.user
	}
	username, oku := gLocalStorage.Get("username")
	token, okt := gLocalStorage.Get("token")
	if !oku || !okt {
		return nil
	}
	d.user = &userData{username, token}
	return d.user
}

func (d *siteData) SetUser(user biz.User) {
	d.user = &userData{user.Username, user.Token}
	gLocalStorage.Set("username", user.Username)
	gLocalStorage.Set("token", user.Token)
}

type LocalStorage struct {
	js.Object
}

func InitLocalStorage() *LocalStorage {
	return &LocalStorage{js.Global.Get("localStorage")}
}

func (ls *LocalStorage) Get(key string) (v string, ok bool) {
	jsv := ls.Object.Call("getItem", key)
	ok = !jsv.IsNull()
	v = jsv.Str()
	return
}

func (ls *LocalStorage) Set(key string, value string) {
	ls.Object.Call("setItem", key, value)
}

func serverUrl(p string) string {
	return path.Join(ServerAddress, p)
}

func main() {
	app := ng.NewModule("publicApp", []string{
		"ngCookies",
		"ngResource",
		"ngSanitize",
		"ngRoute",
	})
	app.NewController("HomeCtrl", func(scope *ng.Scope) {
		scope.Set("user", gSiteData.User())
	})
	app.NewController("UserCtrl", func(scope *ng.Scope, http *ng.HttpService) {
		http.Get(serverUrl("/auth")).Success(func(user biz.User, status int) {
			gSiteData.SetUser(user)
		})
		http.Get(serverUrl("/api/test")).Success(func(user biz.User, status int) {
			println("FUCK")
		})
	})

	app.Config(func(r *ng.RouteProvider) {
		r.When("/", ng.RouteOptions(
			ng.RouteController{"HomeCtrl"},
			ng.RouteTemplate{"views/home.html"},
		))
		r.When("/user", ng.RouteOptions(
			ng.RouteController{"UserCtrl"},
			ng.RouteTemplate{"views/user.html"},
		))
	})

	ng.AddHttpInterceptor(app, "authInterceptor", func(q *ng.QProvider, rootScope *ng.RootScope) ng.HttpInterceptor {
		return ng.HttpInterceptor{
			OnRequest: func(c *ng.ReqSpec) {
				token := gSiteData.User().Token
				if token != "" {
					c.Headers.Value["AuthToken"] = token
				}
			},
			OnResponse: func(j js.Object) interface{} {
				return q.NowOrLater(j)
			},
		}
	})
}
