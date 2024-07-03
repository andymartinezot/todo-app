package models

import "go.mongodb.org/mongo-driver/bson/primitive" //driver used to connect with mongodb

type ToDoList struct{
	ID		primitive.ObjectID		`json:"_id,omitempty" bson:"_id,omitempty"`
	Task	string					`json:"tast,omitempty"`
	Status	bool					`json:"status,omitempty"`
}