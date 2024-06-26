package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/greeflas/racing-engine-backend/internal/model"
	"github.com/greeflas/racing-engine-backend/internal/repository"
)

type RaceArgs struct {
	Id                uuid.UUID `json:"id" validate:"required,uuid"`
	Title             string    `json:"title" validate:"required,max=255"`
	Description       string    `json:"description" validate:"max=2000"`
	ParticipantsCount int       `json:"participants_count" validate:"min=10,max=40"`
	RegistrationAt    time.Time `json:"registration_at"`
	StartAt           time.Time `json:"start_at" validate:"gtefield=RegistrationAt"`
}

type RacesListItem struct {
	Id    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

type RaceService struct {
	raceRepository *repository.RaceRepository
}

type RaceDetailedView struct {
	Id                uuid.UUID `json:"id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	ParticipantsCount int       `json:"participants_count"`
	RegistrationAt    time.Time `json:"registration_at"`
	StartAt           time.Time `json:"start_at"`
}

func NewRaceService(raceRepository *repository.RaceRepository) *RaceService {
	return &RaceService{raceRepository: raceRepository}
}

func (s *RaceService) GetById(ctx context.Context, id uuid.UUID) (*RaceDetailedView, error) {
	race, err := s.raceRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	raceDetailedView := &RaceDetailedView{
		race.Id,
		race.Title,
		race.Description,
		race.ParticipantsCount,
		race.RegistrationAt,
		race.StartAt,
	}

	return raceDetailedView, nil
}

func (s *RaceService) GetList(ctx context.Context) ([]*RacesListItem, error) {
	races, err := s.raceRepository.GetList(ctx)
	if err != nil {
		return nil, err
	}

	racesList := make([]*RacesListItem, len(races))

	for i, race := range races {
		raceItem := RacesListItem{Id: race.Id, Title: race.Title}
		raceItemPointer := &raceItem
		racesList[i] = raceItemPointer
	}

	return racesList, nil
}

func (s *RaceService) Create(ctx context.Context, args *RaceArgs) error {
	race := model.NewRace(
		args.Id,
		args.Title,
		args.Description,
		args.ParticipantsCount,
		args.RegistrationAt,
		args.StartAt,
	)

	return s.raceRepository.Create(ctx, race)
}

func (s *RaceService) Update(ctx context.Context, id uuid.UUID, args *RaceArgs) error {
	race, err := s.raceRepository.GetById(ctx, id)
	if err != nil {
		return err
	}

	race.Update(args.Title, args.Description, args.ParticipantsCount, args.RegistrationAt, args.StartAt)

	return s.raceRepository.Update(ctx, race)
}
