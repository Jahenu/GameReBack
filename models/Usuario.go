package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Usuario struct {
	ID            bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre        string        `bson:"nombre" json:"nombre"`
	Rol           string        `json:"rol" bson:"rol"`
	Correo        string        `bson:"correo" json:"correo"`
	Password      string        `bson:"password" json:"password"`
	FechaRegistro string        `bson:"fechaRegistro" json:"fechaRegistro"`
}
