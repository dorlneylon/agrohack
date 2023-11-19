package points

import (
	`dorlneylon/agrofront/internal/entity`
)

type Points struct {
	Timestamp int      `json:"timestamp"`
	Illness   []string `json:"illness"`
	Data      []Point  `json:"data"`
}

type Point struct {
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Name int     `json:"name"`
}

func ToEntity(dto Points) entity.Points {
	data := make([]entity.Point, len(dto.Data))
	for i, p := range dto.Data {
		data[i] = entity.Point{
			X:    p.X,
			Y:    p.Y,
			Name: dto.Illness[p.Name],
		}
	}
	return entity.Points{
		Timestamp: dto.Timestamp,
		Data:      data,
	}
}

//func FromEntity(entity entity.Points) Points {
//	data := make([]Point, len(entity.Data))
//	for i, p := range entity.Data {
//		data[i] = Point{
//			X:    p.X,
//			Y:    p.Y,
//			Name: p.Name,
//		}
//	}
//	return Points{
//		Timestamp: entity.Timestamp,
//		Data:      data,
//	}
//}
