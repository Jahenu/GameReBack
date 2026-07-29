package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
)

type UsuarioHandler struct {
	UsuarioService *services.UsuarioService
}

func NewUsuarioHandler(service *services.UsuarioService) *UsuarioHandler {
	return &UsuarioHandler{
		UsuarioService: service,
	}
}

// ===============================
// Crear Usuario
// ===============================
func (h *UsuarioHandler) CreateUsuario(c *fiber.Ctx) error {

	var usuario models.Usuario

	if err := c.BodyParser(&usuario); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	if err := h.UsuarioService.Create(&usuario); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Usuario creado correctamente",
		"usuario": usuario,
	})
}

// ===============================
// Login
// ===============================
func (h *UsuarioHandler) Login(c *fiber.Ctx) error {

	var request struct {
		Correo   string `json:"correo"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	usuario, err := h.UsuarioService.Login(
		request.Correo,
		request.Password,
	)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Inicio de sesión correcto",
		"usuario": usuario,
	})
}

// ===============================
// Obtener todos
// ===============================
func (h *UsuarioHandler) GetUsuarios(c *fiber.Ctx) error {

	usuarios, err := h.UsuarioService.FindAll()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(usuarios)
}

// ===============================
// Obtener por ID
// ===============================
func (h *UsuarioHandler) GetUsuarioByID(c *fiber.Ctx) error {

	id := c.Params("id")

	usuario, err := h.UsuarioService.FindByID(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(usuario)
}

// ===============================
// Actualizar
// ===============================
func (h *UsuarioHandler) UpdateUsuario(c *fiber.Ctx) error {

	id := c.Params("id")

	var usuario models.Usuario

	if err := c.BodyParser(&usuario); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	if err := h.UsuarioService.Update(id, &usuario); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Usuario actualizado correctamente",
	})
}

// ===============================
// Eliminar
// ===============================
func (h *UsuarioHandler) DeleteUsuario(c *fiber.Ctx) error {

	id := c.Params("id")

	if err := h.UsuarioService.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Usuario eliminado correctamente",
	})
}
