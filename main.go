package main

import (
	"log"

	"gamerentapi/config"
	"gamerentapi/handlers"
	"gamerentapi/repositories"
	"gamerentapi/routes"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	// Conectar a MongoDB
	if err := config.ConnectMongo(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(cors.New())
	// ==========================
	// Usuario
	// ==========================
	usuarioCollection := config.Database.Collection("usuarios")
	usuarioRepository := repositories.NewUsuarioRepository(usuarioCollection)
	usuarioService := services.NewUsuarioService(usuarioRepository)
	usuarioHandler := handlers.NewUsuarioHandler(usuarioService)

	// ==========================
	// Categoría
	// ==========================
	categoriaCollection := config.Database.Collection("categorias")
	categoriaRepository := repositories.NewCategoriaRepository(categoriaCollection)
	categoriaService := services.NewCategoriaService(categoriaRepository)
	categoriaHandler := handlers.NewCategoriaHandler(categoriaService)

	// ==========================
	// Videojuego
	// ==========================
	videojuegoCollection := config.Database.Collection("videojuegos")
	videojuegoRepository := repositories.NewVideojuegoRepository(videojuegoCollection)
	videojuegoService := services.NewVideojuegoService(videojuegoRepository)
	videojuegoHandler := handlers.NewVideojuegoHandler(videojuegoService)

	// ==========================
	// Renta
	// ==========================
	rentaCollection := config.Database.Collection("rentas")
	rentaRepository := repositories.NewRentaRepository(rentaCollection)
	//rentaService := services.NewRentaService(rentaRepository)
	rentaService := services.NewRentaService(
		rentaRepository,
		videojuegoRepository,
	)
	rentaHandler := handlers.NewRentaHandler(rentaService)

	// ==========================
	// Perfil
	// ==========================
	perfilCollection := config.Database.Collection("perfiles")
	perfilRepository := repositories.NewPerfilRepository(perfilCollection)
	perfilService := services.NewPerfilService(perfilRepository)
	perfilHandler := handlers.NewPerfilHandler(perfilService)

	// ==========================
	// Registrar rutas
	// ==========================
	routes.SetupRoutes(
		app,
		usuarioHandler,
		categoriaHandler,
		videojuegoHandler,
		rentaHandler,
		perfilHandler,
	)

	log.Println("Servidor iniciado en http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
