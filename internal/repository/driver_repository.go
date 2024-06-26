package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/greeflas/racing-engine-backend/pkg/apperror"

	"github.com/greeflas/racing-engine-backend/internal/model"
	"github.com/jackc/pgx/v5"
)

type DriverRepository struct {
	conn *pgx.Conn
}

func NewDriverRepository(conn *pgx.Conn) *DriverRepository {
	return &DriverRepository{conn: conn}
}

func (r *DriverRepository) Create(ctx context.Context, driver *model.Driver) error {
	query := `
INSERT INTO drivers (id, name, phone_number, age, weight, category)
VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.conn.Exec(ctx, query,
		driver.Id,
		driver.Name,
		driver.PhoneNumber,
		driver.Age,
		driver.Weight,
		driver.Category,
	)

	return err
}

func (r *DriverRepository) GetById(ctx context.Context, id uuid.UUID) (*model.Driver, error) {
	query := `SELECT * FROM drivers WHERE id = $1;`

	row := r.conn.QueryRow(ctx, query, id)

	driver := new(model.Driver)

	if err := row.Scan(
		&driver.Id,
		&driver.Name,
		&driver.PhoneNumber,
		&driver.Age,
		&driver.Weight,
		&driver.Category,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperror.NewEntityNotFoundError("Entity driver with id " + id.String() + " not found")
		}

		return nil, err
	}

	return driver, nil
}

func (r *DriverRepository) Update(ctx context.Context, driver *model.Driver) error {
	query := `UPDATE drivers 
		SET (name, phone_number, age, weight, category) = ($1, $2, $3, $4, $5) 	
		WHERE id = $6;`

	_, err := r.conn.Exec(ctx, query,
		driver.Name,
		driver.PhoneNumber,
		driver.Age,
		driver.Weight,
		driver.Category,
		driver.Id,
	)

	return err
}

func (r *DriverRepository) GetList(ctx context.Context) ([]*model.Driver, error) {
	query := `SELECT * FROM drivers`

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var driversList []*model.Driver

	for rows.Next() {
		driver := new(model.Driver)

		err := rows.Scan(
			&driver.Id,
			&driver.Name,
			&driver.PhoneNumber,
			&driver.Age,
			&driver.Weight,
			&driver.Category,
		)
		if err != nil {
			return nil, err
		}

		driversList = append(driversList, driver)
	}

	return driversList, nil
}
