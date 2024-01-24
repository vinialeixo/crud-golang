package repository

import (
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/vinialeixo/crud-golang/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_CreateUser(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("MONGODB_USER_DB", collection_name)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	//defer mtestDb.Close()

	mtestDb.Run("when sending a valid domain returns success", func(mt *mtest.T) {
		email := gofakeit.Email()
		name := gofakeit.Name()
		//resposta de sucesso do banco de dados
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1}, //identifica se a chama foi realizada com sucesso
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)

		userDomain, err := repo.CreteUser(model.NewUserDomain(
			email, "teste", name, 18))

		_, errId := primitive.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.GetEmail(), email)
		assert.EqualValues(t, userDomain.GetName(), name)
	})
}
