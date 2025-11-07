package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)


func (a *application) routes() *chi.Mux {
	// middleware must come before any routes
	a.use(a.Middleware.CheckRemember)

	// add routes here
	a.get("/", a.Handlers.Home)

	a.App.Routes.Get("/users/login", a.Handlers.UserLogin)
	a.post("/users/login", a.Handlers.PostUserLogin)
	a.App.Routes.Get("/users/logout", a.Handlers.Logout)
	a.get("/users/forgot-password", a.Handlers.Forgot)
	a.post("/users/forgot-password", a.Handlers.PostForgot)
	a.get("/users/reset-password", a.Handlers.ResetPasswordForm)
	a.post("/users/reset-password", a.Handlers.PostResetPassword)


	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
