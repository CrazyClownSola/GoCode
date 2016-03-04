package net

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
	"github/sola/repository"
)

type DefaultRoute struct {
	Name         string
	Method       string // RestFul的方法
	Pattern      string
	HandlerFunc  martini.Handler
	ErrorHandler martini.Handler
}

type BindingRoute struct {
	DefaultRoute
	Binding martini.Handler
}

type DefaultRoutes []DefaultRoute
type BindingRoutes []BindingRoute

var get_routes = DefaultRoutes{}

var post_routes = BindingRoutes{
	BindingRoute{
		DefaultRoute{
			"POST_WITH_PARAM",
			"Post",
			"/login",
			Login,
			binding.ErrorHandler,
		},
		binding.Bind(repository.LoginRequest{}),
	},
}

var put_routes = BindingRoutes{}

var del_routes = BindingRoutes{}

func Init_Router() *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin"},
	}))
	m.Use(martini.Static("templates"))
	m.Use(render.Renderer(render.Options{}))

	for _, param := range get_routes {
		m.Get(
			param.Pattern,
			param.HandlerFunc,
			param.ErrorHandler).Name(param.Name)
	}

	for _, param := range post_routes {
		m.Post(
			param.Pattern,
			param.HandlerFunc,
			param.ErrorHandler,
			param.Binding).Name(param.Name)
	}

	for _, param := range put_routes {
		m.Put(
			param.Pattern,
			param.HandlerFunc,
			param.ErrorHandler,
			param.Binding).Name(param.Name)
	}

	for _, param := range del_routes {
		m.Delete(
			param.Pattern,
			param.HandlerFunc,
			param.ErrorHandler,
			param.Binding).Name(param.Name)
	}

	m.Post(
		"/form_test",
		FormTest,
		binding.ErrorHandler,
	).Name("Form_Test")

	return m
}
