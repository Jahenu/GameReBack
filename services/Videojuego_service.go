package services

import (
	"gamerentapi/models"
	"gamerentapi/repositories"
)

type VideojuegoService struct {
	Repository *repositories.VideojuegoRepository
}

func NewVideojuegoService(repository *repositories.VideojuegoRepository) *VideojuegoService {
	return &VideojuegoService{
		Repository: repository,
	}
}

func (s *VideojuegoService) Create(videojuego *models.Videojuego) error {
	return s.Repository.Create(videojuego)
}

func (s *VideojuegoService) FindAll() ([]models.Videojuego, error) {
	return s.Repository.FindAll()
}

func (s *VideojuegoService) FindByID(id string) (*models.Videojuego, error) {
	return s.Repository.FindByID(id)
}

func (s *VideojuegoService) Update(id string, videojuego *models.Videojuego) error {
	return s.Repository.Update(id, videojuego)
}

func (s *VideojuegoService) Delete(id string) error {
	return s.Repository.Delete(id)
}
