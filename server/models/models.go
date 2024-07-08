package models

import "go.mongodb.org/mongo-driver/bson/primitive" //driver used to connect with mongodb

type ToDoList struct{
	ID		primitive.ObjectID		`json:"_id,omitempty" bson:"_id,omitempty"`
	Task	string					`json:"task,omitempty"`
	Status	bool					`json:"status,omitempty"`
}

// type ToDoList struct {
// 	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Task   string             `json:"task" bson:"task"`
// 	Status bool               `json:"status" bson:"status"`
//   }