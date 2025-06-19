package router

import (
	"pembelian-tiket-bioskop-api/internal/controller"
	"pembelian-tiket-bioskop-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                      *fiber.App
	AuthController           controller.AuthController
	FilmController           controller.FilmController
	StudioController         controller.StudioController
	ShowtimeController       controller.ShowtimeController
	AuthenticationMiddleware middleware.Authentication
	AuthorizationMiddleware  middleware.Authorization
}

func (r *RouteConfig) Setup() {
	r.SetupPublicRoute()
	r.SetupAdminRoute()
}

func (r *RouteConfig) SetupPublicRoute() {
	r.App.Post("/api/v1/login", r.AuthController.Login)
	r.App.Post("/api/v1/register", r.AuthController.Register)

	r.App.Get("/api/v1/showtime", r.ShowtimeController.GetShowtimes)
	r.App.Get("/api/v1/showtime/:id", r.ShowtimeController.GetShowtimeByID)

}

func (r *RouteConfig) SetupAdminRoute() {
	admin := r.App.Group("/api/v1/admin", r.AuthenticationMiddleware.Authorize, r.AuthorizationMiddleware.AuthorizeAdmin)

	// film
	film := admin.Group("/film")
	film.Post("/", r.FilmController.CreateFilm)
	film.Get("/", r.FilmController.GetFilms)
	film.Delete("/:id", r.FilmController.DeleteFilm)

	// studio
	studio := admin.Group("/studio")
	studio.Post("/", r.StudioController.CreateStudio)
	studio.Get("/", r.StudioController.GetStudios)
	studio.Delete("/:id", r.StudioController.DeleteStudio)

	// showtime
	showtime := admin.Group("/showtime")
	showtime.Post("/", r.ShowtimeController.CreateShowtime)
	showtime.Patch("/:id", r.ShowtimeController.UpdateShowtime)
	showtime.Delete("/:id", r.ShowtimeController.DeleteShowtime)
}
