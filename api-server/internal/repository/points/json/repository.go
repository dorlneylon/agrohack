package json

import (
	`context`
	`encoding/json`
	`fmt`
	`log/slog`
	`os`
	`path/filepath`
	`time`

	`dorlneylon/agrofront/internal/entity`
	`dorlneylon/agrofront/internal/repository/points`
)

type Repository struct {
	log    *slog.Logger
	path   string
	prefix string
}

func New(log *slog.Logger, path, prefix string) *Repository {
	return &Repository{
		log:    log,
		path:   path,
		prefix: prefix,
	}
}

func (r *Repository) GetAllByDateTime(ctx context.Context, date time.Time) (entity.Points, error) {
	filename := filepath.Join(r.path, fmt.Sprintf("%s_%s.json", r.prefix, date.Format("02_01_2006_15_04")))
	data, err := os.ReadFile(filename)
	if err != nil {
		return entity.Points{}, err
	}

	var dto points.Points
	err = json.Unmarshal(data, &dto)
	if err != nil {
		return entity.Points{}, err
	}

	return points.ToEntity(dto), nil
}
