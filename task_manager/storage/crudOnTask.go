package Storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	cnfig "task_manager/config"
	"errors"
	"fmt"
)

type Storage interface{
	CreateTask(data *models.Task) error
	ReadTask(id string) (models.Task, error) 
	ReadAllTask() ([]models.Task, error)
	UpdateTask(id string, data models.Task) error
	DeleteTask(id string) error
	CreateUser(data *models.User) error
	ReadUser(username string) (models.User, bool) 
	ChangeRoleToAdmin(username string) error
	UpdateUser(username string, data *models.User) error
	NumberOfUsers() (int64, error)
}

type NOSQLConnection struct {
	userCollection *mongo.Collection
	taskCollection *mongo.Collection
}

func NewNoSqlConnection() *NOSQLConnection {
	url := "mongodb://localhost:27017"
	client, err := cnfig.NewMongoStorage(url)
	
	if err != nil{
		fmt.Println(err)
		return nil
	}

	NewTaskcollection := client.Database("task_manager_db").Collection("tasks")
	NewUsercollection := client.Database("task_manager_db").Collection("Users")

	return &NOSQLConnection{
		userCollection : NewUsercollection,
		taskCollection : NewTaskcollection,
	}
}

func (s *NOSQLConnection) CreateTask(data *models.Task) error {
    _, err := s.taskCollection.InsertOne(context.TODO(), data)
    return err
}

func (s *NOSQLConnection) ReadTask(id string) (models.Task, error) {
    var result models.Task
	objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return models.Task{}, fmt.Errorf("invalid object ID: %v", err)
    }

    err = s.taskCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&result)

    if err != nil {
        if err == mongo.ErrNoDocuments{
            return models.Task{}, errors.New("document not found")
        }
        return models.Task{}, err
    }
    return result, nil
}

func (s *NOSQLConnection) ReadAllTask() ([]models.Task, error) {
    var results []models.Task
    cursor, err := s.taskCollection.Find(context.TODO(), bson.M{})

    if err != nil {
        return nil, err
    }

    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var result models.Task
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

func (s *NOSQLConnection) UpdateTask(id string, data models.Task) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
        return err
    }
	filter := bson.M{"_id" : objectID}
	update :=  bson.M{"$set": data}
	_, err = s.taskCollection.UpdateOne(context.Background(), filter, update)
	return err
}


func (s *NOSQLConnection) DeleteTask(id string) error{
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
        return err
    }
	filter := bson.M{"_id" : objectID}
	result, err := s.taskCollection.DeleteOne(context.TODO(), filter)
	
	if err != nil{
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("document with ID %s not found", id)
	}

	return nil
}