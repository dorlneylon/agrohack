package getAll

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	`time`

	`dorlneylon/agrofront/internal/api/http/request`
	`dorlneylon/agrofront/internal/api/http/response`
	`dorlneylon/agrofront/internal/entity`
	`dorlneylon/agrofront/internal/lib/logger/sl`
	`dorlneylon/agrofront/internal/usecase/points`
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Request struct {
	Date string `json:"date"`
}

type Response struct {
	entity.Points
}

type Getter interface {
	GetAllByDateTime(ctx context.Context, date time.Time) (entity.Points, error)
}

func New(log *slog.Logger, gt Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "api.http.link.add"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		req, err := request.PrepareRequest[Request](r)
		if err != nil {
			log.Error("invalid request", sl.Err(err))
			response.Error(w, r, err, 400)
			return
		}

		date, err := time.Parse("02.01.2006 15:04:05", req.Date)
		if err != nil {
			log.Error("fail to parse date", sl.Err(err))
			response.Error(w, r, err, 400)
			return
		}

		points, err := gt.GetAllByDateTime(context.Background(), date)
		if err != nil {
			handleGetPointsError(log, w, r, err)
			return
		}

		render.JSON(w, r, Response{
			Points: points,
		})
	}
}

func handleGetPointsError(log *slog.Logger, w http.ResponseWriter, r *http.Request, err error) {
	log.Error("failed to get points", sl.Err(err))
	if errors.Is(err, points.ErrFailedToGetPoints) {
		response.Error(w, r, err, 404)
	} else {
		response.InternalError(w, r)
	}
}
