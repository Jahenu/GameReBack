package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Renta struct {
	ID           bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UsuarioID    bson.ObjectID `bson:"usuarioId" json:"usuarioId"`
	Titulo       string        `bson:"titulo" json:"titulo"`
	VideojuegoID bson.ObjectID `bson:"videojuegoId" json:"videojuegoId"`
	FechaRenta   string        `bson:"fechaRenta" json:"fechaRenta"`
	FechaEntrega string        `bson:"fechaEntrega" json:"fechaEntrega"`
	Estado       string        `bson:"estado" json:"estado"`
}
