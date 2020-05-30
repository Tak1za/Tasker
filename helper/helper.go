package helper

import (
	"github.com/Tak1za/tasker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTasks(payload []primitive.M, detailed bool) (interface{}, error) {
	if detailed {
		var results []models.ToDoList
		for _, d := range payload {
			res, err := convertToTask(d)
			if err != nil {
				return nil, err
			}
			results = append(results, res)
		}
		return results, nil
	} else {
		var results []string
		for _, d := range payload {
			res, err := convertToTask(d)
			if err != nil {
				return nil, err
			}
			results = append(results, res.Task)
		}
		return results, nil
	}
}

func GetTask(payload primitive.M, detailed bool) (interface{}, error) {
	res, err := convertToTask(payload)
	if err != nil {
		return nil, err
	}
	if detailed {
		return res, nil
	} else {
		return res.Task, nil
	}
}

func convertToTask(data primitive.M) (models.ToDoList, error) {
	var task models.ToDoListDB
	b, _ := bson.Marshal(data)
	err := bson.Unmarshal(b, &task)
	if err != nil {
		return models.ToDoList{}, err
	}
	res := models.ToDoList{
		ID:     task.ID.Hex(),
		Task:   task.Task,
		Status: task.Status,
	}
	return res, nil
}
