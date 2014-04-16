package main

import (
	"encoding/json"
	"fmt"
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
			fmt.Println(err)
		} else {
			defer res.Body.Close()

			err = json.NewDecoder(res.Body).Decode(&v)

			if err != nil {
				fmt.Println(err)
			}
		}

		r.HTML(200, "new", v)
	})

	m.Run()
}
