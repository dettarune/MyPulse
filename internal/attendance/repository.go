package attendance

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, employeeID int64, lat, lng float64) (Attendance, error) {
	var a Attendance
	err := r.db.QueryRow(ctx,
		`INSERT INTO attendances (employee_id, latitude, longitude)
         VALUES ($1, $2, $3)
         RETURNING id, employee_id, checked_in_at, latitude, longitude`,
		employeeID, lat, lng,
	).Scan(&a.ID, &a.EmployeeID, &a.CheckedInAt, &a.Latitude, &a.Longitude)
	if err != nil {
		return Attendance{}, fmt.Errorf("simpan absensi: %w", err)
	}
	return a, nil
}
