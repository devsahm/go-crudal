package services

import (
	"context"
	"crude-app/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.userCollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	err := u.userCollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {

	var users []*models.User

	cursor, err := u.userCollection.Find(u.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("No  record was found")
	}

	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {

	query := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$Set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_age", Value: user.Age}, bson.E{Key: "user_address", Value: user.Address}}}}

	result, _ := u.userCollection.UpdateOne(u.ctx, query, update)

	if result.MatchedCount != 1 {
		return errors.New("no match record was found for update")
	}

	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {

	query := bson.D{bson.E{Key: "user_name", Value: name}}
	result, _ := u.userCollection.DeleteOne(u.ctx, query)

	if result.DeletedCount != 1 {
		return errors.New("No match was found for this user name")
	}

	return nil
}
