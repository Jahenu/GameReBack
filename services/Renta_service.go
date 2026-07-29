/*package services

import (
	"errors"
	"gamerentapi/models"
	"gamerentapi/repositories"
)

type RentaService struct {
	Repository *repositories.RentaRepository
}

// NUEVO MÉTODO
func (s *RentaService) FindByUsuario(usuarioID string) ([]models.Renta, error) {
	return s.Repository.FindByUsuario(usuarioID)
}

func NewRentaService(repository *repositories.RentaRepository) *RentaService {
	return &RentaService{
		Repository: repository,
	}
}

func (s *RentaService) Create(renta *models.Renta) error {

	existe, err := s.Repository.ExisteRentaActiva(
		renta.UsuarioID,
		renta.VideojuegoID,
	)

	if err != nil {
		return err
	}

	if existe {
		return errors.New("este videojuego ya está rentado")
	}

	return s.Repository.Create(renta)
}

func (s *RentaService) FindAll() ([]models.Renta, error) {
	return s.Repository.FindAll()
}

func (s *RentaService) FindByID(id string) (*models.Renta, error) {
	return s.Repository.FindByID(id)
}

func (s *RentaService) Update(id string, renta *models.Renta) error {
	return s.Repository.Update(id, renta)
}

func (s *RentaService) Delete(id string) error {
	return s.Repository.Delete(id)
}
*/

package services

import (
	"errors"

	"gamerentapi/models"
	"gamerentapi/repositories"
)

type RentaService struct {
	Repository           *repositories.RentaRepository
	VideojuegoRepository *repositories.VideojuegoRepository
}

func NewRentaService(
	repository *repositories.RentaRepository,
	videojuegoRepository *repositories.VideojuegoRepository,
) *RentaService {

	return &RentaService{
		Repository:           repository,
		VideojuegoRepository: videojuegoRepository,
	}
}

func (s *RentaService) Create(renta *models.Renta) error {

	existe, err := s.Repository.ExisteRentaActiva(
		renta.UsuarioID,
		renta.VideojuegoID,
	)

	if err != nil {
		return err
	}

	if existe {
		return errors.New("este videojuego ya está rentado")
	}

	videojuego, err := s.VideojuegoRepository.FindByID(
		renta.VideojuegoID.Hex(),
	)

	if err != nil {
		return err
	}

	if videojuego.Stock <= 0 {
		return errors.New("no hay stock disponible")
	}

	videojuego.Stock--

	err = s.VideojuegoRepository.Update(
		videojuego.ID.Hex(),
		videojuego,
	)

	if err != nil {
		return err
	}

	return s.Repository.Create(renta)
}

func (s *RentaService) FindAll() ([]models.Renta, error) {
	return s.Repository.FindAll()
}

func (s *RentaService) FindByID(id string) (*models.Renta, error) {
	return s.Repository.FindByID(id)
}

func (s *RentaService) Update(id string, renta *models.Renta) error {
	return s.Repository.Update(id, renta)
}

func (s *RentaService) Delete(id string) error {
	return s.Repository.Delete(id)
}

func (s *RentaService) FindByUsuario(usuarioID string) ([]models.Renta, error) {
	return s.Repository.FindByUsuario(usuarioID)
}
