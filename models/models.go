package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDoListDB struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `form:"task" json:"task" binding:"required"`
	Status bool               `json:"status,omitempty"`
}

type ToDoList struct {
	ID     string `json:"id,omitempty"`
	Task   string `json:"task"`
	Status bool   `json:"status,omitempty"`
}

func (t ToDoList) String() {
	fmt.Println("ID: ", t.ID)
	fmt.Println("Task: ", t.Task)
	fmt.Println("Status: ", t.Status)
	fmt.Print("\n")
}
