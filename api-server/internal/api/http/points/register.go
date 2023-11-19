package points

import (
	"context"
	"log/slog"
	`time`

	`dorlneylon/agrofront/internal/api/http/points/getAll`
	`dorlneylon/agrofront/internal/entity`
	"github.com/go-chi/chi/v5"
)

type usecase interface {
	GetAllByDateTime(ctx context.Context, date time.Time) (entity.Points, error)
}

func Register(r *chi.Mux, pointsUsecase usecase, log *slog.Logger) {
	r.Post("/points/getAll", getAll.New(log, pointsUsecase))
}
