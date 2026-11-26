package attendance

import (
	"context"
	"errors"
)

var ErrInvalidLocation = errors.New("lokasi tidak valid")

// store didefinisikan di sisi pemakai supaya service mudah dites dengan mock.
type store interface {
	Create(ctx context.Context, employeeID int64, lat, lng float64) (Attendance, error)
}

type Service struct {
	store store
}

func NewService(s store) *Service {
	return &Service{store: s}
}

func (s *Service) CheckIn(ctx context.Context, employeeID int64, lat, lng float64) (Attendance, error) {
	if lat < -90 || lat > 90 || lng < -180 || lng > 180 {
		return Attendance{}, ErrInvalidLocation
	}
	return s.store.Create(ctx, employeeID, lat, lng)
}
