package repositories

import (
	"context"
	"gamerentapi/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RentaRepository struct {
	Collection *mongo.Collection
}

func NewRentaRepository(collection *mongo.Collection) *RentaRepository {
	return &RentaRepository{
		Collection: collection,
	}
}

func (r *RentaRepository) Create(renta *models.Renta) error {

	_, err := r.Collection.InsertOne(
		context.Background(),
		renta,
	)

	return err
}

func (r *RentaRepository) FindAll() ([]models.Renta, error) {

	cursor, err := r.Collection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return nil, err
	}

	var rentas []models.Renta

	err = cursor.All(
		context.Background(),
		&rentas,
	)

	return rentas, err
}

func (r *RentaRepository) FindByID(id string) (*models.Renta, error) {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var renta models.Renta

	err = r.Collection.FindOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
	).Decode(&renta)

	if err != nil {
		return nil, err
	}

	return &renta, nil
}

func (r *RentaRepository) Update(id string, renta *models.Renta) error {

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
				"usuario_id":    renta.UsuarioID,
				"videojuego_id": renta.VideojuegoID,
				"titulo":        renta.Titulo,
				"fecha_renta":   renta.FechaRenta,
				"fecha_entrega": renta.FechaEntrega,
				"estado":        renta.Estado,
			},
		},
	)

	return err
}

func (r *RentaRepository) Delete(id string) error {

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

func (r *RentaRepository) ExisteRentaActiva(usuarioID, videojuegoID bson.ObjectID) (bool, error) {

	var renta models.Renta

	err := r.Collection.FindOne(
		context.Background(),
		bson.M{
			"usuarioId":    usuarioID,
			"videojuegoId": videojuegoID,
			"estado":       "Activa",
		},
	).Decode(&renta)

	if err == mongo.ErrNoDocuments {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *RentaRepository) FindByUsuario(usuarioID string) ([]models.Renta, error) {

	objectID, err := bson.ObjectIDFromHex(usuarioID)

	if err != nil {
		return nil, err
	}

	cursor, err := r.Collection.Find(
		context.Background(),
		bson.M{
			"usuarioId": objectID,
		},
	)

	if err != nil {
		return nil, err
	}

	var rentas []models.Renta

	err = cursor.All(
		context.Background(),
		&rentas,
	)

	if err != nil {
		return nil, err
	}

	return rentas, nil
}
