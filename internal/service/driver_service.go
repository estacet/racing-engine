package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/greeflas/racing-engine-backend/internal/model"
	"github.com/greeflas/racing-engine-backend/internal/repository"
)

type CreateDriverArgs struct {
	Id          uuid.UUID `json:"id" validate:"required,uuid"`
	Name        string    `json:"name" validate:"required,min=2,max=100"`
	PhoneNumber string    `json:"phone_number" validate:"required,phoneNumber"`
	Age         *int      `json:"age" validate:"min=12,max=99"`
	Weight      *int      `json:"weight" validate:"min=40,max=85"`
}

type UpdateDriverArgs struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	PhoneNumber string `json:"phone_number" validate:"required,phoneNumber"`
	Age         *int   `json:"age" validate:"min=12,max=99"`
	Weight      *int   `json:"weight" validate:"min=40,max=85"`
}

type DriversListItem struct {
	Id       uuid.UUID      `json:"id"`
	Name     string         `json:"name"`
	Category model.Category `json:"category"`
}

type DriverDetailedInfo struct {
	Id          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	PhoneNumber string         `json:"phone_number"`
	Age         *int           `json:"age"`
	Weight      *int           `json:"weight"`
	Category    model.Category `json:"category"`
}

type DriverService struct {
	driverRepository *repository.DriverRepository
}

func NewDriverService(driverRepository *repository.DriverRepository) *DriverService {
	return &DriverService{driverRepository: driverRepository}
}

func (s *DriverService) Create(ctx context.Context, args *CreateDriverArgs) error {
	driver := model.NewDriver(
		args.Id,
		args.Name,
		args.PhoneNumber,
		args.Age,
		args.Weight,
	)

	return s.driverRepository.Create(ctx, driver)
}

func (s *DriverService) Update(ctx context.Context, id uuid.UUID, args *UpdateDriverArgs) error {
	driver, err := s.driverRepository.GetById(ctx, id)
	if err != nil {
		return err
	}

	driver.Update(args.Name, args.PhoneNumber, args.Age, args.Weight)

	return s.driverRepository.Update(ctx, driver)
}

func (s *DriverService) GetList(ctx context.Context) ([]*DriversListItem, error) {
	drivers, err := s.driverRepository.GetList(ctx)
	if err != nil {
		return nil, err
	}

	driversList := make([]*DriversListItem, len(drivers))

	for i, driver := range drivers {
		driverItem := &DriversListItem{
			Id:       driver.Id,
			Name:     driver.Name,
			Category: driver.Category,
		}
		driversList[i] = driverItem
	}

	return driversList, nil
}

func (s *DriverService) GetById(ctx context.Context, id uuid.UUID) (*DriverDetailedInfo, error) {
	driver, err := s.driverRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	driverDetailedInfo := &DriverDetailedInfo{
		driver.Id,
		driver.Name,
		driver.PhoneNumber,
		driver.Age,
		driver.Weight,
		driver.Category,
	}

	return driverDetailedInfo, nil
}
