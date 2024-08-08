package mongodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/task_manager/Domain"
	"errors"
	"context"
)

type MongoUserRepo struct{
	collection *mongo.Collection
}


func NewMongoUserRepo() *MongoUserRepo {
	url := "mongodb://localhost:27017"
	client, err := NewMongoStorage(url)

	if err != nil{
		fmt.Println(err)
		return nil
	}

	NewUsercollection := client.Database("task_manager_db").Collection("Users")

	return &MongoUserRepo{
		collection : NewUsercollection,
	}
}

func (s *MongoUserRepo) CreateUser(data *domain.User) error{
	_, err := s.collection.InsertOne(context.TODO(), data)
	return err
}

func (s *MongoUserRepo) NumberOfUsers() (int64, error){
	return s.collection.CountDocuments(context.TODO(), bson.D{})
}

func (s *MongoUserRepo) ReadUser(username string) (domain.User, bool) {
	var result domain.User
	err := s.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&result)

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
	update :=  bson.M{"$set": data}
	_, err := s.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (s *MongoUserRepo) ChangeRoleToAdmin(username string) error{
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"role": "Admin"}}

	result, err := s.collection.UpdateOne(context.TODO(), filter, update)
	if result.MatchedCount == 0{
		return errors.New("the user doesb not exist")
    	}

    	return err
}
