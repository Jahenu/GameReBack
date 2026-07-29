package routes

import (
	"gamerentapi/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	usuarioHandler *handlers.UsuarioHandler,
	categoriaHandler *handlers.CategoriaHandler,

	videojuegoHandler *handlers.VideojuegoHandler,
	rentaHandler *handlers.RentaHandler,
	perfilHandler *handlers.PerfilHandler,
) {

	// ==========================
	// Ruta principal
	// ==========================
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"mensaje": "GameRent API funcionando",
		})
	})

	// ==========================
	// Login
	// ==========================
	app.Post("/login", usuarioHandler.Login)

	// ==========================
	// Usuarios
	// ==========================
	app.Get("/usuarios", usuarioHandler.GetUsuarios)
	app.Post("/usuarios", usuarioHandler.CreateUsuario)
	app.Get("/usuarios/:id", usuarioHandler.GetUsuarioByID)
	app.Put("/usuarios/:id", usuarioHandler.UpdateUsuario)
	app.Delete("/usuarios/:id", usuarioHandler.DeleteUsuario)

	// ==========================
	// Rentas
	// ==========================
	app.Get("/rentas", rentaHandler.GetRentas)
	app.Get("/rentas/usuarios/:id", rentaHandler.GetRentasByUsuario)
	app.Post("/rentas", rentaHandler.CreateRenta)
	app.Get("/rentas/:id", rentaHandler.GetRenta)
	app.Put("/rentas/:id", rentaHandler.UpdateRenta)
	app.Delete("/rentas/:id", rentaHandler.DeleteRenta)

	// ==========================
	// Categorías
	// ==========================
	app.Get("/categorias", categoriaHandler.GetCategorias)
	app.Post("/categorias", categoriaHandler.CreateCategoria)
	app.Get("/categorias/:id", categoriaHandler.GetCategoria)
	app.Put("/categorias/:id", categoriaHandler.UpdateCategoria)
	app.Delete("/categorias/:id", categoriaHandler.DeleteCategoria)

	// ==========================
	// Perfiles
	// ==========================
	app.Get("/perfiles", perfilHandler.GetPerfiles)
	app.Post("/perfiles", perfilHandler.CreatePerfil)
	app.Get("/perfiles/:id", perfilHandler.GetPerfil)
	app.Put("/perfiles/:id", perfilHandler.UpdatePerfil)
	app.Delete("/perfiles/:id", perfilHandler.DeletePerfil)

	// ==========================
	// Videojuegos
	// ==========================
	app.Get("/videojuego", videojuegoHandler.GetVideojuegos)
	app.Post("/videojuego", videojuegoHandler.CreateVideojuego)
	app.Get("/videojuego/:id", videojuegoHandler.GetVideojuego)
	app.Put("/videojuego/:id", videojuegoHandler.UpdateVideojuego)
	app.Delete("/videojuego/:id", videojuegoHandler.DeleteVideojuego)

}
