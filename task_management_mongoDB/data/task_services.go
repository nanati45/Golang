package data

import (
	"context"
	"errors"
	"fmt"
	"task_management_mongoDB/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "tasks"

// MongoDB client and collection initialization
var taskCollection *mongo.Collection

func InitMongoDB(client *mongo.Client, dbName string) {
	taskCollection = client.Database(dbName).Collection(collectionName)
}

// Get all tasks
func GetAllTasks(ctx context.Context) ([]models.Task, error) {
	cursor, err := taskCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// Get a task by ID
func GetTaskByID(ctx context.Context, id string) (models.Task, error) {
	var task models.Task
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return task, fmt.Errorf("invalid ID format: %v", err)
	}
	filter := bson.M{"_id": objectID}
	err = taskCollection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, err
	}
	return task, nil
}

// Add a task
func AddTask(ctx context.Context, task models.Task) error {
	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	_, err := taskCollection.InsertOne(ctx, task)
	return err
}

// Update a task
func UpdateTask(ctx context.Context, id string, updatedTask models.Task) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %v", err)
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"title":       updatedTask.Title,
		"description": updatedTask.Description,
		"dueDate":     updatedTask.DueDate,
		"status":      updatedTask.Status,
		"updatedAt":   time.Now(),
	}}
	result, err := taskCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}


// Remove a task
func RemoveTask(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %v", err)
	}
	filter := bson.M{"_id": objectID}
	result, err := taskCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

