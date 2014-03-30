package main

import ng "github.com/gopherjs/go-angularjs"

func main() {
	app := ng.NewModule("publicApp", []string{
		"ngCookies",
		"ngResource",
		"ngSanitize",
		"ngRoute",
	})
	app.NewController("MainCtrl", func(scope *ng.Scope) {
		scope.Set("awesomeThings", []string{"fuck", "eat"})
	})
	app.Config(func(services *ng.Injector) {
		ng.RouteProvider(services).When("/", ng.Options(
			ng.Route.Controller("MainCtrl"),
			ng.Route.TemplatePath("views/main.html"),
		))
	}, ng.Inject("$routeProvider"))
	println("Hello from gopherjs")
}
