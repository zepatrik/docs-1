package actions

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/buffalo/render/resolvers"
	"github.com/gobuffalo/gobuffalo/actions/helpers"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		HTMLLayout:     "application.html",
		CacheTemplates: ENV == "production",
		Helpers: map[string]interface{}{
			"panel": helpers.PanelHelper,
		},
		FileResolverFunc: func() resolvers.FileResolver {
			return &resolvers.RiceBox{Box: rice.MustFindBox("../templates")}
		},
	})
}

func assetsPath() http.FileSystem {
	if ENV == "production" {
		return http.Dir("/app/assets")
	}

	box := rice.MustFindBox("../assets")
	return box.HTTPBox()
}
