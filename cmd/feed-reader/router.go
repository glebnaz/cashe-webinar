package main

import (
	"github.com/glebnaz/cache-webinar/internal/router"
	"net/http"
)

var Router = []router.Router{
	{
		Path:    "/v1/create-post",
		Method:  http.MethodPost,
		Handler: server.NewCreatePost(),
	},
}
