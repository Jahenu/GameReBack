package repositories

import (
	"context"
	"gamerentapi/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type VideojuegoRepository struct {
	Collection *mongo.Collection
}

func NewVideojuegoRepository(collection *mongo.Collection) *VideojuegoRepository {
	return &VideojuegoRepository{
		Collection: collection,
	}
}

func (r *VideojuegoRepository) Create(videojuego *models.Videojuego) error {
	_, err := r.Collection.InsertOne(
		context.Background(),
		videojuego,
	)

	return err
}

func (r *VideojuegoRepository) FindAll() ([]models.Videojuego, error) {

	cursor, err := r.Collection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	videojuegos := make([]models.Videojuego, 0)

	if err = cursor.All(
		context.Background(),
		&videojuegos,
	); err != nil {
		return nil, err
	}

	return videojuegos, nil
}

func (r *VideojuegoRepository) FindByID(id string) (*models.Videojuego, error) {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var videojuego models.Videojuego

	err = r.Collection.FindOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
	).Decode(&videojuego)

	if err != nil {
		return nil, err
	}

	return &videojuego, nil
}

func (r *VideojuegoRepository) Update(
	id string,
	videojuego *models.Videojuego,
) error {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = r.Collection.UpdateOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
		bson.M{
			"$set": bson.M{
				"titulo":      videojuego.Titulo,
				"plataforma":  videojuego.Plataforma,
				"precioRenta": videojuego.PrecioRenta,
				"stock":       videojuego.Stock,
				"activo":      videojuego.Activo,
				"categoriaId": videojuego.CategoriaID,
			},
		},
	)

	return err
}

func (r *VideojuegoRepository) Delete(id string) error {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = r.Collection.DeleteOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
	)

	return err
}
