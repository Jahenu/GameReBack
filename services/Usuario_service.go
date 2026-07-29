/*package services

import (
	"errors"
	"time"

	"gamerentapi/models"
	"gamerentapi/repositories"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UsuarioService struct {
	UsuarioRepository *repositories.UsuarioRepository
}

func NewUsuarioService(repo *repositories.UsuarioRepository) *UsuarioService {
	return &UsuarioService{
		UsuarioRepository: repo,
	}
}

func (s *UsuarioService) Create(usuario *models.Usuario) error {

	if usuario.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}

	if usuario.Correo == "" {
		return errors.New("el correo es obligatorio")
	}

	if usuario.Password == "" {
		return errors.New("la contraseña es obligatoria")
	}

	return s.UsuarioRepository.Create(usuario)
}

func (s *UsuarioService) Login(correo, password string) (*models.Usuario, error) {

	usuario, err := s.UsuarioRepository.FindByCorreo(correo)

	if err != nil {

		// Si no existe el usuario, lo crea automáticamente
		if err == mongo.ErrNoDocuments {

			nuevoUsuario := &models.Usuario{
				Nombre:        correo,
				Correo:        correo,
				Password:      password,
				FechaRegistro: time.Now().Format("2006-01-02"),
			}

			if err := s.UsuarioRepository.Create(nuevoUsuario); err != nil {
				return nil, err
			}

			return s.UsuarioRepository.FindByCorreo(correo)
		}

		return nil, err
	}

	// Validar contraseña
	if usuario.Password != password {
		return nil, errors.New("contraseña incorrecta")
	}

	return usuario, nil
}

func (s *UsuarioService) FindAll() ([]*models.Usuario, error) {
	return s.UsuarioRepository.FindAll()
}

func (s *UsuarioService) FindByID(id string) (*models.Usuario, error) {
	return s.UsuarioRepository.FindByID(id)
}

func (s *UsuarioService) Update(id string, usuario *models.Usuario) error {
	return s.UsuarioRepository.Update(id, usuario)
}

func (s *UsuarioService) Delete(id string) error {
	return s.UsuarioRepository.Delete(id)
}
*/

package services

import (
	"errors"
	"time"

	"gamerentapi/models"
	"gamerentapi/repositories"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UsuarioService struct {
	UsuarioRepository *repositories.UsuarioRepository
}

func NewUsuarioService(repo *repositories.UsuarioRepository) *UsuarioService {
	return &UsuarioService{
		UsuarioRepository: repo,
	}
}

func (s *UsuarioService) Create(usuario *models.Usuario) error {

	if usuario.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}

	if usuario.Correo == "" {
		return errors.New("el correo es obligatorio")
	}

	if usuario.Password == "" {
		return errors.New("la contraseña es obligatoria")
	}

	// Si no se especifica un rol, se asigna "cliente"
	if usuario.Rol == "" {
		usuario.Rol = "cliente"
	}

	return s.UsuarioRepository.Create(usuario)
}

func (s *UsuarioService) Login(correo, password string) (*models.Usuario, error) {

	usuario, err := s.UsuarioRepository.FindByCorreo(correo)

	if err != nil {

		// Si no existe el usuario, lo crea automáticamente
		if err == mongo.ErrNoDocuments {

			nuevoUsuario := &models.Usuario{
				Nombre:        correo,
				Rol:           "cliente",
				Correo:        correo,
				Password:      password,
				FechaRegistro: time.Now().Format("2006-01-02"),
			}

			if err := s.UsuarioRepository.Create(nuevoUsuario); err != nil {
				return nil, err
			}

			return s.UsuarioRepository.FindByCorreo(correo)
		}

		return nil, err
	}

	// Validar contraseña
	if usuario.Password != password {
		return nil, errors.New("contraseña incorrecta")
	}

	return usuario, nil
}

func (s *UsuarioService) FindAll() ([]*models.Usuario, error) {
	return s.UsuarioRepository.FindAll()
}

func (s *UsuarioService) FindByID(id string) (*models.Usuario, error) {
	return s.UsuarioRepository.FindByID(id)
}

func (s *UsuarioService) Update(id string, usuario *models.Usuario) error {
	return s.UsuarioRepository.Update(id, usuario)
}

func (s *UsuarioService) Delete(id string) error {
	return s.UsuarioRepository.Delete(id)
}
