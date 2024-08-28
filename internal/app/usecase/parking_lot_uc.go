package usecase

import (
	"errors"

	"github.com/CamiloLeonP/parking-radar/internal/app/domain"
	"github.com/CamiloLeonP/parking-radar/internal/app/repository"
)

type IParkingLotUseCase interface {
	CreateParkingLot(req CreateParkingLotRequest) error
	GetParkingLot(parkingLotID uint) (*domain.ParkingLot, error)
	UpdateParkingLot(parkingLotID uint, req UpdateParkingLotRequest) error
	DeleteParkingLot(parkingLotID uint) error
	ListParkingLots() ([]domain.ParkingLot, error)
}

type ParkingLotUseCase struct {
	ParkingLotRepository repository.IParkingLotRepository
}

type CreateParkingLotRequest struct {
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	TotalSpaces int     `json:"total_spaces"`
}

type UpdateParkingLotRequest struct {
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	TotalSpaces int     `json:"total_spaces"`
}

func NewParkingLotUseCase(parkingLotRepo repository.IParkingLotRepository) IParkingLotUseCase {
	return &ParkingLotUseCase{
		ParkingLotRepository: parkingLotRepo,
	}
}

func (uc *ParkingLotUseCase) CreateParkingLot(req CreateParkingLotRequest) error {
	parkingLot := domain.ParkingLot{
		Name:        req.Name,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		TotalSpaces: req.TotalSpaces,
	}

	return uc.ParkingLotRepository.Create(&parkingLot)
}

func (uc *ParkingLotUseCase) GetParkingLot(parkingLotID uint) (*domain.ParkingLot, error) {
	return uc.ParkingLotRepository.GetByID(parkingLotID)
}

func (uc *ParkingLotUseCase) UpdateParkingLot(parkingLotID uint, req UpdateParkingLotRequest) error {
	parkingLot, err := uc.ParkingLotRepository.GetByID(parkingLotID)
	if err != nil {
		return err
	}

	parkingLot.Name = req.Name
	parkingLot.Latitude = req.Latitude
	parkingLot.Longitude = req.Longitude

	// Actualizar availableSpaces solo si totalSpaces cambia
	if parkingLot.TotalSpaces != req.TotalSpaces {
		difference := req.TotalSpaces - parkingLot.TotalSpaces
		parkingLot.AvailableSpaces += difference

		if parkingLot.AvailableSpaces < 0 {
			return errors.New(("no available spaces left to occupy"))
		}
	}

	parkingLot.TotalSpaces = req.TotalSpaces

	return uc.ParkingLotRepository.Update(parkingLot)
}

func (uc *ParkingLotUseCase) DeleteParkingLot(parkingLotID uint) error {
	return uc.ParkingLotRepository.Delete(parkingLotID)
}

func (uc *ParkingLotUseCase) ListParkingLots() ([]domain.ParkingLot, error) {
	return uc.ParkingLotRepository.List()
}
