package main

import (
	ng "github.com/gopherjs/go-angularjs"
	jscon "honnef.co/go/js/console"
)

func main() {
	app := ng.NewModule("publicApp", []string{
		"ngCookies",
		"ngResource",
		"ngSanitize",
		"ngRoute",
	})
	mainCtrl := app.NewController("MainCtrl")
	app.Config(func(r ng.RouteProvider) {
		r.When("/", ng.RouteConfig{
			TemplateUrl: "views/main.html",
			Controller:  "MainCtrl",
		}).
			Otherwise(ng.RouteConfig{
			RedirectTo: "/",
		})
	}, "$routeProvider")
	jscon.Log("Hello from gopherjs")
}
