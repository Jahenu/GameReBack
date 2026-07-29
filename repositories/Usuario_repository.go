package repositories

import (
	"context"

	"gamerentapi/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UsuarioRepository struct {
	Collection *mongo.Collection
}

func NewUsuarioRepository(collection *mongo.Collection) *UsuarioRepository {
	return &UsuarioRepository{
		Collection: collection,
	}
}

func (r *UsuarioRepository) Create(usuario *models.Usuario) error {

	_, err := r.Collection.InsertOne(
		context.Background(),
		usuario,
	)

	return err
}

func (r *UsuarioRepository) FindAll() ([]*models.Usuario, error) {

	cursor, err := r.Collection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var usuarios []*models.Usuario

	if err = cursor.All(context.Background(), &usuarios); err != nil {
		return nil, err
	}

	return usuarios, nil
}

func (r *UsuarioRepository) FindByID(id string) (*models.Usuario, error) {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var usuario models.Usuario

	err = r.Collection.FindOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
	).Decode(&usuario)

	if err != nil {
		return nil, err
	}

	return &usuario, nil
}

func (r *UsuarioRepository) FindByCorreo(correo string) (*models.Usuario, error) {

	var usuario models.Usuario

	err := r.Collection.FindOne(
		context.Background(),
		bson.M{
			"correo": correo,
		},
	).Decode(&usuario)

	if err != nil {
		return nil, err
	}

	return &usuario, nil
}

func (r *UsuarioRepository) Update(id string, usuario *models.Usuario) error {

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
				"nombre":        usuario.Nombre,
				"rol":           usuario.Rol,
				"correo":        usuario.Correo,
				"password":      usuario.Password,
				"fechaRegistro": usuario.FechaRegistro,
			},
		},
	)

	return err
}

func (r *UsuarioRepository) Delete(id string) error {

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
