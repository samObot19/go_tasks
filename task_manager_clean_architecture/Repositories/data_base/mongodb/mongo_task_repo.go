package mongodb

import (
	"github.com/task_manager/Domain"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"errors"
	"fmt"
)

type MongoTaskRepo struct{
	collection *mongo.Collection
}


func NewMongoTaskRepo() *NOSQLConnection {
	url := "mongodb://localhost:27017"
	client, err := NewMongoStorage(url)

	if err != nil{
		fmt.Println(err)
		return nil
	}

	NewTaskcollection := client.Database("task_manager_db").Collection("tasks")
	

	return &MongoTaskRepo{
		collection : NewTaskcollection,
	}
}

func (s *MongoTaskRepo) CreateTask(data *domain.Task) error {
	_, err := s.collection.InsertOne(context.TODO(), data)
	return err
}

func (s *MongoTaskRepo) ReadTask(id string) (domain.Task, error) {
	var result domain.Task
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return domain.Task{}, fmt.Errorf("invalid object ID: %v", err)
    	}

    	err = s.collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&result)

    	if err != nil {
        	if err == mongo.ErrNoDocuments{
           		return domain.Task{}, errors.New("document not found")
        	}
        	return domain.Task{}, err
    	}
    	return result, nil
}

func (s *MongoTaskReop) ReadAllTask() ([]domain.Task, error) {
	var results []domain.Task
	cursor, err := s.collection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

    	defer cursor.Close(context.TODO())

    	for cursor.Next(context.TODO()) {
        	var result domain.Task
        	if err := cursor.Decode(&result); err != nil {
            		return nil , err
        	}
        	results = append(results, result)
    	}

    	if err := cursor.Err(); err != nil {
        	return nil, err
    	}

    	return results, nil
}

func (s *MongoTaskRepo) UpdateTask(id string, data domain.Task) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
        	return err
    	}
	filter := bson.M{"_id" : objectID}
	update :=  bson.M{"$set": data}
	_, err = s.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (s *MongoTaskRepo) DeleteTask(id string) error{
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
        	return err
    	}
	filter := bson.M{"_id" : objectID}
	result, err := s.collection.DeleteOne(context.TODO(), filter)

	if err != nil{
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("document with ID %s not found", id)
	}

	return nil
}
