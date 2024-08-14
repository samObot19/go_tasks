package mongodb_test

import (
	"context"
	"testing"
	"time"
	"os"
	"github.com/stretchr/testify/suite"
	"github.com/task_manager/Domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/task_manager/Repositories/data_base/mongodb"
	
)

type MongoTaskRepoTestSuite struct {
	suite.Suite
	repo       *mongodb.MongoTaskRepo
	collection *mongo.Collection
	client     *mongo.Client
}



func (suite * MongoTaskRepoTestSuite) SetupSuite() {
	// Connect to MongoDB
	url :=  os.Getenv("url")
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	suite.NoError(err)
	suite.client = client

	// Set up the test collection
	suite.collection = client.Database("task_manager_test_db").Collection("tasks")
	suite.repo = &mongodb.MongoTaskRepo{
		Collection: suite.collection,
		Client:     client,
    	}
}

	
func (suite *MongoTaskRepoTestSuite) TearDownSuite() {
	// Drop the test database after all tests are done
	err := suite.client.Database("task_manager_test_db").Drop(context.Background())
	suite.NoError(err)

	// Disconnect the client
	err = suite.client.Disconnect(context.Background())
	suite.NoError(err)
}

func (suite *MongoTaskRepoTestSuite) SetupTest() {
	// Clear the collection before each test
	err := suite.collection.Drop(context.Background())
	suite.NoError(err)
}

func (suite *MongoTaskRepoTestSuite) TestCreateTask() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}

	err := suite.repo.CreateTask(task)
	suite.NoError(err)

	// Verify the task was inserted
	var insertedTask domain.Task
	err = suite.collection.FindOne(context.Background(), primitive.M{"title": "Test Task"}).Decode(&insertedTask)
	suite.NoError(err)
	suite.Equal(task.Title, insertedTask.Title)
	suite.Equal(task.Description, insertedTask.Description)
	suite.Equal(task.Status, insertedTask.Status)
}


func (suite *MongoTaskRepoTestSuite) TestReadTask() {
	// Insert a task into the collection directly
	task := domain.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}
	insertResult, err := suite.collection.InsertOne(context.Background(), task)
	suite.NoError(err)
	taskID := insertResult.InsertedID.(primitive.ObjectID).Hex()

	// Read the task using the repository
	readTask, err := suite.repo.ReadTask(taskID)
	suite.NoError(err)
	suite.Equal(task.Title, readTask.Title)
	suite.Equal(task.Description, readTask.Description)
	suite.Equal(task.Status, readTask.Status)
}

func (suite *MongoTaskRepoTestSuite) TestReadAllTask() {
	// Insert tasks into the collection directly
	tasks := []domain.Task{
		{
			Title:       "Test Task 1",
			Description: "This is the first test task",
			DueDate:     time.Now().Add(24 * time.Hour),
			Status:      "pending",
		},
		{
			Title:       "Test Task 2",
			Description: "This is the second test task",
			DueDate:     time.Now().Add(48 * time.Hour),
			Status:      "completed",
		},
	}
	 taskInterfaces := make([]interface{}, len(tasks))
	 for i, task := range tasks {
		 taskInterfaces[i] = task

	 }
	_, err := suite.collection.InsertMany(context.Background(), taskInterfaces)
	suite.NoError(err)

	// Read all tasks using the repository
	readTasks, err := suite.repo.ReadAllTask()
	suite.NoError(err)
	suite.Equal(2, len(readTasks))
}

func (suite *MongoTaskRepoTestSuite) TestUpdateTask() {
	// Insert a task into the collection directly
	task := domain.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}
	insertResult, err := suite.collection.InsertOne(context.Background(), task)
	suite.NoError(err)
	taskID := insertResult.InsertedID.(primitive.ObjectID).Hex()

	// Update the task using the repository
	task.Title = "Updated Task"
	task.Status = "completed"
	err = suite.repo.UpdateTask(taskID, task)
	suite.NoError(err)

	// Verify the task was updated
	var updatedTask domain.Task
	err = suite.collection.FindOne(context.Background(), primitive.M{"_id": insertResult.InsertedID}).Decode(&updatedTask)
	suite.NoError(err)
	suite.Equal(task.Title, updatedTask.Title)
	suite.Equal(task.Status, updatedTask.Status)
}

func (suite *MongoTaskRepoTestSuite) TestDeleteTask() {
	// Insert a task into the collection directly
	task := domain.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}
	insertResult, err := suite.collection.InsertOne(context.Background(), task)
	suite.NoError(err)
	taskID := insertResult.InsertedID.(primitive.ObjectID).Hex()

	// Delete the task using the repository
	err = suite.repo.DeleteTask(taskID)
	suite.NoError(err)

	// Verify the task was deleted
	count, err := suite.collection.CountDocuments(context.Background(), primitive.M{"_id": insertResult.InsertedID})
	suite.NoError(err)
	suite.Equal(int64(0), count)
}

func TestMongoTaskRepoTestSuite(t *testing.T) {
	suite.Run(t, new(MongoTaskRepoTestSuite))
}
