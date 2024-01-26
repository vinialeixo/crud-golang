package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/vinialeixo/crud-golang/src/model/repository"
	"github.com/vinialeixo/crud-golang/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("MONGODB_USER_DB", collection_name)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	email := gofakeit.Email()
	name := gofakeit.Name()
	userEntity := entity.UserEntity{
		ID:       primitive.NewObjectID(),
		Email:    email,
		Password: "teste",
		Name:     name,
		Age:      18,
	}
	mtestDb.Run("when sending a valid email returns success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch, convertEntityToBson(userEntity)))

		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)
		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomain.GetName(), userDomain.GetName())
	})

	mtestDb.Run("when error when mongodb returns error", func(mt *mtest.T) {
		//forçar um error, não para mockar uma mensagem de sucesso

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByEmail("teste")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)

	})

	mtestDb.Run("returns no documents found", func(mt *mtest.T) {
		//forçar um error, não para mockar uma mensagem de sucesso

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name), mtest.FirstBatch))
		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByEmail("teste")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)

	})
}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}

func TestUserRepository_FindUserByID(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("MONGODB_USER_DB", collection_name)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	email := gofakeit.Email()
	name := gofakeit.Name()
	userEntity := entity.UserEntity{
		ID:       primitive.NewObjectID(),
		Email:    email,
		Password: "teste",
		Name:     name,
		Age:      18,
	}
	mtestDb.Run("when sending a valid id returns success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch, convertEntityToBson(userEntity)))

		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByID(userEntity.ID.Hex())
		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomain.GetName(), userDomain.GetName())
	})

	mtestDb.Run("when error when mongodb returns error", func(mt *mtest.T) {
		//forçar um error, não para mockar uma mensagem de sucesso

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByID("teste")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)

	})

	mtestDb.Run("returns no documents found", func(mt *mtest.T) {
		//forçar um error, não para mockar uma mensagem de sucesso

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name), mtest.FirstBatch))
		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByID("teste")

		assert.NotNil(t, err.Message, fmt.Sprintf("User not found with this ID: teste "))
		assert.Nil(t, userDomain)

	})
}

func TestUserRepository_FindUserByEmailAndPassowrd(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("MONGODB_USER_DB", collection_name)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	email := gofakeit.Email()
	name := gofakeit.Name()
	userEntity := entity.UserEntity{
		ID:       primitive.NewObjectID(),
		Email:    email,
		Password: "teste",
		Name:     name,
		Age:      18,
	}
	mtestDb.Run("when sending a valid email and password returns success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch, convertEntityToBson(userEntity)))

		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword(userEntity.Email, userEntity.Password)
		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomain.GetName(), userDomain.GetName())
	})

	mtestDb.Run("when error when mongodb returns error", func(mt *mtest.T) {
		//forçar um error, não para mockar uma mensagem de sucesso

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByEmailAndPassword("teste", "testepass")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)

	})

	mtestDb.Run("returns no documents found", func(mt *mtest.T) {
		//forçar um error, não para mockar uma mensagem de sucesso

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name), mtest.FirstBatch))
		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByEmailAndPassword("teste", "testepass")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)

	})
}
