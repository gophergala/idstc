package main

import (
	"net/http"

	"github.com/unrolled/render"
	"idstc/model"
)

func (c *ctx) Home(w http.ResponseWriter, r *http.Request) {
	ValidateRoute("", w, r)

	p := model.Page{
		"dashboard",
		"dashboard",
		nil,
	}

	c.r.HTML(w, http.StatusOK, "home/index", p)
}

func (c *ctx) People(w http.ResponseWriter, r *http.Request) {
	ValidateRoute("people", w, r)

	p := model.Page{
		"people",
		"people",
		nil,
	}

	c.r.HTML(w, http.StatusOK, "home/people", p)
}

func (c *ctx) Projects(w http.ResponseWriter, r *http.Request) {
	ValidateRoute("projects", w, r)

	p := model.Page{
		"projects",
		"projects",
		nil,
	}

	c.r.HTML(w, http.StatusOK, "home/projects", p)
}

func (c *ctx) Analytics(w http.ResponseWriter, r *http.Request) {
	ValidateRoute("analytics", w, r)

	p := model.Page{
		"analytics",
		"analytics",
		nil,
	}

	c.r.HTML(w, http.StatusOK, "home/analytics", p)
}

func (c *ctx) Account(w http.ResponseWriter, r *http.Request) {
	ValidateRoute("account", w, r)

	p := model.Page{
		"account",
		"account",
		nil,
	}

	c.r.HTML(w, http.StatusOK, "home/account", p)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	ren := render.New(render.Options{
		Layout: "shared/layout",
	})

	p := model.Page{
		"Page not found",
		"404",
		nil,
	}

	ren.HTML(w, http.StatusOK, "404", p)
}