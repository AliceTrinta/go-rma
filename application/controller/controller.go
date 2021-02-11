package controller

import (
	"GoWonder/handler"
	"GoWonder/route"
	"net/http"
)

//Here, at the Begin Function, the routes are configured.
func Begin() http.Handler {

	routes := []route.Conf{
		{Path: "/", F: handler.Index, Method: "GET"},
	}

	return route.Route(routes)
}
