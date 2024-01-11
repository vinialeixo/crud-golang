package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` //se o campo for vazio ignorar ele, criando um usuario não terá um usuario
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Name     string             `bson:"name"`
	Age      int8               `bson:"age"`
}
