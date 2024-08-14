package mongodb_test

import (
	"context"
	"testing"
	//"fmt"
	"time"
	"os"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/task_manager/Domain"
	"github.com/task_manager/Repositories/data_base/mongodb"
)

type MongoUserRepoTestSuite struct {
	suite.Suite
	repo       *mongodb.MongoUserRepo
	collection *mongo.Collection
	client     *mongo.Client
}

func (suite *MongoUserRepoTestSuite) SetupSuite() {
	// Connect to MongoDB
	url :=  os.Getenv("url")
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	suite.NoError(err)
	suite.client = client

	// Set up the test collection
	suite.collection = client.Database("task_manager_test_db").Collection("Users")
	suite.repo = &mongodb.MongoUserRepo{
		Collection: suite.collection,
		Client : suite.client,
	}
}

func (suite *MongoUserRepoTestSuite) TearDownSuite() {
	// Drop the test database after the tests are complete
	err := suite.client.Database("task_manager_test_db").Drop(context.Background())
	suite.NoError(err)
	
	
	// Disconnect the client
	err = suite.client.Disconnect(context.Background())
	suite.NoError(err)
}

func (suite *MongoUserRepoTestSuite) SetupTest() {
    _, err := suite.collection.DeleteMany(context.Background(), bson.D{})
    if err != nil {
        suite.FailNow("Failed to clear the collection before test", err)
    }
}
func (suite *MongoUserRepoTestSuite) TestCreateUser() {
	user := &domain.User{
		UserName: "testuser",
		Password: "password123",
		Role:     "User",
	}

	err := suite.repo.CreateUser(user)
	suite.NoError(err)
	time.Sleep(100 * time.Millisecond)

	// Verify the user was created
	var result domain.User
	err = suite.collection.FindOne(context.Background(), bson.M{"username": "testuser"}).Decode(&result)
	suite.NoError(err)
	suite.Equal("testuser", result.UserName)
	suite.Equal("password123", result.Password)
	suite.Equal("User", result.Role)
}

func (suite *MongoUserRepoTestSuite) TestNumberOfUsers() {
	 user := &domain.User{
                UserName: "testuser",
                Password: "password123",
                Role:     "User",
        }
	err := suite.repo.CreateUser(user)
	suite.NoError(err)

	count, err := suite.repo.NumberOfUsers()
	suite.NoError(err)
	suite.Equal(int64(1), count)
}

func (suite *MongoUserRepoTestSuite) TestReadUser() {
	 User := &domain.User{
                UserName: "testuser",
                Password: "password123",
                Role:     "User",
        }
        err := suite.repo.CreateUser(User)

        suite.NoError(err)
	user, found := suite.repo.ReadUser("testuser")
	suite.True(found)
	suite.Equal("testuser", user.UserName)
	suite.Equal("password123", user.Password)
}

func (suite *MongoUserRepoTestSuite) TestUpdateUser() {
    	user := &domain.User{
        	UserName: "testuser",
       		Password: "password123",
        	Role:     "User",
    	}

    	err := suite.repo.CreateUser(user)
    	suite.NoError(err)

    	// Update only the password
    	updatedUser := &domain.User{
        	Password: "newpassword123",
    	}

    	err = suite.repo.UpdateUser("testuser", updatedUser)
    	suite.NoError(err)

    // Verify the user was updated
    	var result domain.User
    	err = suite.collection.FindOne(context.Background(), bson.M{"username": "testuser"}).Decode(&result)
    	suite.NoError(err)
    	suite.Equal("newpassword123", result.Password)
    	suite.Equal("User", result.Role)  // Ensure the role is still "User"
}

func (suite *MongoUserRepoTestSuite) TestChangeRoleToAdmin() {
	// Insert a user into the collection
	user := &domain.User{
		UserName: "testuser",
		Password: "password123",
		Role:     "User",
	}
	err := suite.repo.CreateUser(user)
	suite.NoError(err, "should not return an error when creating user")

	// Verify the user was created
	createdUser, exists := suite.repo.ReadUser("testuser")
	suite.True(exists, "user should exist after creation")
	suite.Equal("User", createdUser.Role, "user role should be 'User' before update")

	// Change the user's role to admin
	err = suite.repo.ChangeRoleToAdmin("testuser")
	suite.NoError(err, "should not return an error when changing role")

	// Verify the role was updated
	updatedUser, exists := suite.repo.ReadUser("testuser")
	suite.True(exists, "user should exist after role update")
	suite.Equal("Admin", updatedUser.Role, "user role should be updated to Admin")
}

func TestMongoUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(MongoUserRepoTestSuite))
}
