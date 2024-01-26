package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinialeixo/crud-golang/src/model/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_DeleteUser(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("MONGODB_USER_DB", collection_name)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	//defer mtestDb.Close()

	mtestDb.Run("when sending a valid userId return success", func(mt *mtest.T) {

		//resposta de sucesso do banco de dados
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1}, //identifica se a chama foi realizada com sucesso
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)

		err := repo.DeleteUser("teste")

		assert.Nil(t, err)

	})

	mtestDb.Run("return error from database", func(mt *mtest.T) {

		//resposta de sucesso do banco de dados
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)
		err := repo.DeleteUser("teste")

		assert.NotNil(t, err)

	})

}
