package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"net/url"
)

func main() {
	m := martini.Classic()

	//middleware
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", "")
	})

	m.Get("/new", func(r render.Render) {
		var v map[string]interface{}

		values := url.Values{}
		res, err := http.PostForm("http://requestb.in/api/v1/bins", values)

		if err != nil {
			r.JSON(500, map[string]interface{}{"error": err})
		} else {
			defer res.Body.Close()

			err = json.NewDecoder(res.Body).Decode(&v)

			if err != nil {
				r.JSON(500, map[string]interface{}{"error": err})
			}
		}

		r.HTML(200, "new", v)
	})

	m.Run()
}
