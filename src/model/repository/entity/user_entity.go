package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` //se o campo for vazio ignorar ele, criando um usuario não terá um usuario
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Age      int8               `bson:"age,omitempty"`
}
