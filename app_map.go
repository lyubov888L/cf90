package main

import (
	"log"
)

func init() {
	register(&Command{
		group: "Application",
		name:  "app.map",
		help:  "Map application to given host and domain (route must already exist)",
		params: []Param{
			Param{name: "name", desc: "Application name"},
			Param{name: "host", desc: "Host name"},
			Param{name: "domain", desc: "Domain name"},
		},
		handle: app_map,
	})
}

func app_map() {

	target, err := c.SelectedTarget()
	if err != nil {
		log.Fatal(err)
	}
	apps, err := target.AppsGet()
	if err != nil {
		log.Fatal(err)
	}

	// Get Application ID
	name, appId := params["name"], ""
	for _, app := range apps {
		if app.Name == name {
			appId = app.Guid
		}
	}
	if appId == "" {
		i, err := choose(AppList(apps))
		if err != nil {
			log.Fatal(err)
		}
		appId = apps[i].Guid
	}

	// Get RouteID
	host := params["host"]
	domain := params["domain"]
	routes, err := target.RoutesGet()
	if err != nil {
		log.Fatal(err)
	}
	routeGUID := ""
	for _, route := range routes {
		if route.Host == host && route.Domain.Name == domain {
			routeGUID = route.Guid
		}
	}
	if routeGUID == "" {
		index, err := choose(RouteList(routes))
		if err != nil {
			log.Fatal(err)
		}
		routeGUID = routes[index].Guid
	}

	// Perform action
	err = target.AppAddRoute(appId, routeGUID)
	if err != nil {
		log.Fatal(err)
	}

}
