package Storage

import (
	"context"
	"EnhancingTaskManagementAPI/Models"
	cnfig "EnhancingTaskManagementAPI/Configration"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"errors"
	"fmt"
)

type Storage interface{
	Create(data interface{}) error
	Read(id string) (interface{}, error) 
	ReadAll() (interface{}, error)
	Update(id string, data interface{}) error
	Delete(id string) error
}

type NOSQLConnection struct {
	Collection *mongo.Collection
}

func NewNoSqlConnection() *NOSQLConnection {
	url := "mongodb://localhost:27017"
	client, err := cnfig.NewMongoStorage(url)
	
	if err != nil{
		fmt.Println(err)
		return nil
	}

	Newcollection := client.Database("task_manager_db").Collection("tasks")
	return &NOSQLConnection{Collection: Newcollection}
}

func (s *NOSQLConnection) Create(data interface{}) error {
    _, err := s.Collection.InsertOne(context.TODO(), data)
    return err
}

func (s *NOSQLConnection) Read(id string) (interface{}, error) {
    var result Models.Task
    err := s.Collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&result)

    if err != nil {
        if err == mongo.ErrNoDocuments{
            return nil, errors.New("document not found")
        }
        return nil, err
    }
    return result, nil
}

func (s *NOSQLConnection) ReadAll() (interface{}, error) {
    var results [] Models.Task
    cursor, err := s.Collection.Find(context.TODO(), bson.M{})

    if err != nil {
        return nil, err
    }

    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var result Models.Task
        if err := cursor.Decode(&result); err != nil {
            return nil, err
        }
        results = append(results, result)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return results, nil
}

func (s *NOSQLConnection) Update(id string, data interface{}) error {
	filter := bson.M{"id": id}
	update :=  bson.M{"$set": data}
	_, err := s.Collection.UpdateOne(context.Background(), filter, update)
	return err
}


func (s *NOSQLConnection) Delete(id string) error{
	filter := bson.M{"id" : id}
	result, err := s.Collection.DeleteOne(context.TODO(), filter)
	
	if err != nil{
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("document with ID %s not found", id)
	}

	return nil
}
