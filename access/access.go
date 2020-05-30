package access

import (
	"context"
	"log"

	"github.com/Tak1za/tasker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://admin:varunak47@cluster0-3v4tr.mongodb.net/test?retryWrites=true&w=majority"
const dbName = "test"
const collectionName = "todolist"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(collectionName)
}

func GetTasks() ([]primitive.M, error) {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())

	var results []primitive.M

	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			return nil, e
		}

		results = append(results, result)
	}

	err = cur.Err()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func GetTask(id string) (primitive.M, error) {
	var result bson.M
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(context.Background(), bson.D{{"_id", objId}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func AddTask(task models.ToDoList) (primitive.ObjectID, error) {
	insertResult, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		return primitive.NilObjectID, err
	}

	log.Println("Inserted record with ID: ", insertResult.InsertedID.(primitive.ObjectID).Hex())
	return insertResult.InsertedID.(primitive.ObjectID), err
}
