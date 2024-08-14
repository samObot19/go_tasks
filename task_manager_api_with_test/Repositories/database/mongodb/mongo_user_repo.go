package mongodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/task_manager/Domain"
	"errors"
	"context"
	"github.com/joho/godotenv"
	"log"
    "os"
)

type MongoUserRepo struct{
	Collection *mongo.Collection
	Client *mongo.Client
}


func NewMongoUserRepo() *MongoUserRepo {
	err := godotenv.Load("../.env")
    if err != nil {
		fmt.Println(err)
        log.Fatalf("Error loading .env file:")
    }

	url := os.Getenv("url")

	client, err := NewMongoStorage(url)

	if err != nil{
		fmt.Println(err)
		return nil
	}

	NewUsercollection := client.Database("task_manager_db").Collection("Users")

	return &MongoUserRepo{
		Collection : NewUsercollection,
		Client : client,
	}
}

func (s *MongoUserRepo) CreateUser(data *domain.User) error{
	_, err := s.Collection.InsertOne(context.TODO(), data)
	return err
}

func (s *MongoUserRepo) NumberOfUsers() (int64, error){
	return s.Collection.CountDocuments(context.TODO(), bson.D{})
}

func (s *MongoUserRepo) ReadUser(username string) (domain.User, bool) {
	var result domain.User
	err := s.Collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&result)

    	if err != nil {
        	if err == mongo.ErrNoDocuments{
            		return domain.User{}, false
        	}
        	return domain.User{}, false
    	}
    	return result, true
}

func (s *MongoUserRepo) UpdateUser(username string, data *domain.User) error {
    filter := bson.M{"username": username}
    update := bson.M{"$set": bson.M{}}

    if data.Password != "" {
        update["$set"].(bson.M)["password"] = data.Password
    }
    if data.Role != "" {
        update["$set"].(bson.M)["role"] = data.Role
    }

    _, err := s.Collection.UpdateOne(context.Background(), filter, update)
    return err
}


func (s *MongoUserRepo) ChangeRoleToAdmin(username string) error {
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"role": "Admin"}}

	result, err := s.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("the user does not exist")
	}

	return nil
}
