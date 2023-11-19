package points

import (
	`context`
	`errors`
	`time`

	`dorlneylon/agrofront/internal/entity`
)

type repository interface {
	GetAllByDateTime(ctx context.Context, date time.Time) (entity.Points, error)
}

type statService interface {
	ProceedAllByDateTime(ctx context.Context, date time.Time) error
}

type Usecase struct {
	r repository
	s statService
}

func New(r repository, s statService) *Usecase {
	return &Usecase{
		r: r,
		s: s,
	}
}

var (
	ErrFailedToGetPoints = errors.New("Error during obtaining points")
)

func (u *Usecase) GetAllByDateTime(ctx context.Context, date time.Time) (entity.Points, error) {
	points, err := u.r.GetAllByDateTime(ctx, date)

	if err != nil {
		//log.Println("usecase.repo err", err)
		err = u.s.ProceedAllByDateTime(ctx, date)
		return entity.Points{}, ErrFailedToGetPoints
		//if err != nil {
		//}
		//return points, nil
	}

	return points, nil
}
