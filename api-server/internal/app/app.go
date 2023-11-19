package app

import (
	"context"
	"log/slog"
	"net/http"

	points_router `dorlneylon/agrofront/internal/api/http/points`
	`dorlneylon/agrofront/internal/config`
	`dorlneylon/agrofront/internal/repository/points/json`
	points_service `dorlneylon/agrofront/internal/services/points`
	points_usecase `dorlneylon/agrofront/internal/usecase/points`
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type App struct {
	log *slog.Logger
	cfg *config.Config
	srv *http.Server
}

func New(log *slog.Logger, cfg *config.Config) *App {
	return &App{
		log: log,
		cfg: cfg,
	}
}

// Start init app and start http server in goroutine
func (a *App) Start() {
	r := a.initRouter()
	a.initPoints(r)
	a.log.Info("Starting server...")
	a.startServer(r)
}

func (a *App) initPoints(r *chi.Mux) {
	pointsStorage := json.New(a.log, a.cfg.JsonStorage.Path, a.cfg.JsonStorage.Prefix)
	service := points_service.New(a.log, a.cfg.Service)
	pointsUsecase := points_usecase.New(pointsStorage, service)
	points_router.Register(r, pointsUsecase, a.log)
}

func (a *App) initUser(r *chi.Mux) {
}

func (a *App) initRouter() *chi.Mux {
	r := chi.NewRouter()
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of the major browsers.
	})
	r.Use(corsMiddleware.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	return r
}

func (a *App) startServer(r *chi.Mux) {
	srv := &http.Server{
		Addr:         a.cfg.HTTPServer.Address,
		Handler:      r,
		ReadTimeout:  a.cfg.HTTPServer.Timeout,
		WriteTimeout: a.cfg.HTTPServer.Timeout,
		IdleTimeout:  a.cfg.HTTPServer.Timeout,
	}
	a.srv = srv

	go srv.ListenAndServe()
}

func (a *App) Stop(ctx context.Context) {
	a.srv.Shutdown(ctx)
}
