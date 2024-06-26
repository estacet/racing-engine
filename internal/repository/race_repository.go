package repository

import (
	"context"
	"errors"

	"github.com/greeflas/racing-engine-backend/pkg/apperror"

	"github.com/google/uuid"

	"github.com/greeflas/racing-engine-backend/internal/model"
	"github.com/jackc/pgx/v5"
)

type RaceRepository struct {
	conn *pgx.Conn
}

func NewRaceRepository(conn *pgx.Conn) *RaceRepository {
	return &RaceRepository{conn: conn}
}

func (r *RaceRepository) GetById(ctx context.Context, id uuid.UUID) (*model.Race, error) {
	query := `SELECT * FROM races WHERE id = $1;`

	row := r.conn.QueryRow(ctx, query, id)

	race := new(model.Race)

	if err := row.Scan(
		&race.Id,
		&race.Title,
		&race.Description,
		&race.ParticipantsCount,
		&race.RegistrationAt,
		&race.StartAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperror.NewEntityNotFoundError("Entity race with id " + id.String() + " not found")
		}

		return nil, err
	}

	return race, nil
}

func (r *RaceRepository) GetList(ctx context.Context) ([]*model.Race, error) {
	query := `SELECT * FROM races`

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var racesList []*model.Race

	for rows.Next() {
		race := new(model.Race)

		err := rows.Scan(
			&race.Id,
			&race.Title,
			&race.Description,
			&race.ParticipantsCount,
			&race.RegistrationAt,
			&race.StartAt,
		)
		if err != nil {
			return nil, err
		}

		racesList = append(racesList, race)
	}

	return racesList, nil
}

func (r *RaceRepository) Create(ctx context.Context, race *model.Race) error {
	query := `
INSERT INTO races (id, title, description, participants_count, registration_at, start_at)
VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.conn.Exec(ctx, query,
		race.Id,
		race.Title,
		race.Description,
		race.ParticipantsCount,
		race.RegistrationAt,
		race.StartAt,
	)

	return err
}

func (r *RaceRepository) Update(ctx context.Context, race *model.Race) error {
	query := `UPDATE races 
		SET (title, description, participants_count, registration_at, start_at) = ($1, $2, $3, $4, $5) 	
		WHERE id = $6;`

	_, err := r.conn.Exec(ctx, query,
		race.Title,
		race.Description,
		race.ParticipantsCount,
		race.RegistrationAt,
		race.StartAt,
		race.Id,
	)

	return err
}
