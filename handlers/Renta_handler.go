package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
)

type RentaHandler struct {
	Service *services.RentaService
}

func NewRentaHandler(service *services.RentaService) *RentaHandler {
	return &RentaHandler{
		Service: service,
	}
}

func (h *RentaHandler) CreateRenta(c *fiber.Ctx) error {

	var renta models.Renta

	if err := c.BodyParser(&renta); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	err := h.Service.Create(&renta)

	if err != nil {

		// Si el servicio detecta que ya existe una renta activa
		if err.Error() == "este videojuego ya está rentado" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Renta registrada correctamente",
		"renta":   renta,
	})
}

func (h *RentaHandler) GetRentas(c *fiber.Ctx) error {

	rentas, err := h.Service.FindAll()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(rentas)
}

func (h *RentaHandler) GetRenta(c *fiber.Ctx) error {

	id := c.Params("id")

	renta, err := h.Service.FindByID(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Renta no encontrada",
		})
	}

	return c.JSON(renta)
}

func (h *RentaHandler) GetRentasByUsuario(c *fiber.Ctx) error {

	usuarioID := c.Params("id")

	rentas, err := h.Service.FindByUsuario(usuarioID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(rentas)
}
func (h *RentaHandler) UpdateRenta(c *fiber.Ctx) error {

	id := c.Params("id")

	var renta models.Renta

	if err := c.BodyParser(&renta); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	if err := h.Service.Update(id, &renta); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Renta actualizada",
	})
}

func (h *RentaHandler) DeleteRenta(c *fiber.Ctx) error {

	id := c.Params("id")

	if err := h.Service.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Renta eliminada",
	})
}
