package config

import (
	"pembelian-tiket-bioskop-api/internal/controller"
	"pembelian-tiket-bioskop-api/internal/middleware"
	"pembelian-tiket-bioskop-api/internal/repository"
	"pembelian-tiket-bioskop-api/internal/router"
	"pembelian-tiket-bioskop-api/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func StartServer(config *AppConfig) {
	authRepository := repository.NewAuthRepository(config.DB)
	filmRepository := repository.NewFilmRepository(config.DB)
	studioRepository := repository.NewStudioRepository(config.DB)
	showtimeRepository := repository.NewShowtimeRepository(config.DB)

	authenticationMiddleware := middleware.NewAuthenticationMiddleware(config.Config)
	authorizationMiddleware := middleware.NewAuthorizationMiddleware(authenticationMiddleware)

	authService := service.NewAuthService(config.Log, config.Config, config.Validate, authRepository, authenticationMiddleware)
	filmService := service.NewFilmService(config.Log, filmRepository, config.Validate)
	studioService := service.NewStudioService(config.Log, studioRepository, config.Validate)
	showtimeService := service.NewShowtimeService(config.Log, showtimeRepository, config.Validate)

	authController := controller.NewAuthController(config.Log, authService)
	filmController := controller.NewFilmController(config.Log, filmService)
	studioController := controller.NewStudioController(config.Log, studioService)
	showtimeController := controller.NewShowtimeController(config.Log, showtimeService)

	routeConfig := router.RouteConfig{
		App:                      config.App,
		AuthController:           authController,
		FilmController:           filmController,
		StudioController:         studioController,
		ShowtimeController:       showtimeController,
		AuthenticationMiddleware: authenticationMiddleware,
		AuthorizationMiddleware:  authorizationMiddleware,
	}

	routeConfig.Setup()
}
