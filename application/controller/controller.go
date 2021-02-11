package controller

import (
	"net/http"

	"github.com/AliceTrinta/GoWonder/handler"
	"github.com/AliceTrinta/GoWonder/route"
)

//Begin Function will configure the routes.
func Begin() http.Handler {

	routes := []route.Conf{
		{Path: "/", F: handler.Index, Method: "GET"},
	}

	return route.Route(routes)
}
