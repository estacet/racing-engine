package model

import (
	"time"

	"github.com/google/uuid"
)

type Race struct {
	Id                uuid.UUID
	Title             string
	Description       string
	ParticipantsCount int
	RegistrationAt    time.Time
	StartAt           time.Time
}

func NewRace(
	id uuid.UUID,
	title string,
	description string,
	participantsCount int,
	registrationAt time.Time,
	startAt time.Time,
) *Race {
	return &Race{
		Id:                id,
		Title:             title,
		Description:       description,
		ParticipantsCount: participantsCount,
		RegistrationAt:    registrationAt,
		StartAt:           startAt,
	}
}

func (r *Race) Update(
	title string,
	description string,
	participantsCount int,
	registrationAt time.Time,
	startAt time.Time,
) {
	r.Title = title
	r.Description = description
	r.ParticipantsCount = participantsCount
	r.RegistrationAt = registrationAt
	r.StartAt = startAt
}
