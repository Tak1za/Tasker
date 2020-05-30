package helper

import (
	"github.com/Tak1za/tasker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTasks(payload []primitive.M, detailed bool) (interface{}, error) {
	var task models.ToDoList

	if detailed {
		var results []models.ToDoList
		var task models.ToDoListDB
		var res models.ToDoList
		for _, d := range payload {
			b, _ := bson.Marshal(d)
			err := bson.Unmarshal(b, &task)
			if err != nil {
				return nil, err
			}

			res.ID = task.ID.Hex()
			res.Task = task.Task
			res.Status = task.Status
			results = append(results, res)
		}
		return results, nil
	} else {
		var results []string
		for _, d := range payload {
			b, _ := bson.Marshal(d)
			err := bson.Unmarshal(b, &task)
			if err != nil {
				return nil, err
			}
			results = append(results, task.Task)
		}

		return results, nil
	}
}

func GetTask(payload primitive.M, detailed bool) (interface{}, error) {
	if detailed {
		var task models.ToDoListDB
		var res models.ToDoList
		b, _ := bson.Marshal(payload)
		err := bson.Unmarshal(b, &task)
		if err != nil {
			return nil, err
		}

		res.ID = task.ID.Hex()
		res.Task = task.Task
		res.Status = task.Status

		return res, nil
	} else {
		var task models.ToDoListDB
		b, _ := bson.Marshal(payload)
		err := bson.Unmarshal(b, &task)
		if err != nil {
			return nil, err
		}

		return task.Task, nil
	}
}
