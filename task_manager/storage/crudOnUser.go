package Storage

import (
	"task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
    "errors"
)

func (s *NOSQLConnection) CreateUser(data *models.User) error{
    _, err := s.userCollection.InsertOne(context.TODO(), data)
    return err
}

func (s *NOSQLConnection) NumberOfUsers() (int64, error){
    return s.userCollection.CountDocuments(context.TODO(), bson.D{})
}

func (s *NOSQLConnection) ReadUser(username string) (models.User, bool) {
    var result models.User
    err := s.userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&result)

    if err != nil {
        if err == mongo.ErrNoDocuments{
            return models.User{}, false
        }
        return models.User{}, false
    }
    return result, true
}


func (s *NOSQLConnection) UpdateUser(username string, data *models.User) error {
	filter := bson.M{"username": username}
	update :=  bson.M{"$set": data}
	_, err := s.userCollection.UpdateOne(context.Background(), filter, update)
	return err
}


func (s *NOSQLConnection) ChangeRoleToAdmin(username string) error{
    filter := bson.M{"username": username}
    update := bson.M{"$set": bson.M{"role": "Admin"}}

    result, err := s.userCollection.UpdateOne(context.TODO(), filter, update)

    if result.MatchedCount == 0{
        return errors.New("the user doesb not exist")
    }

    return err
}