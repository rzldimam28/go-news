package res

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewsResponse struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title,omitempty" bson:"title,omitempty"`
	Content string             `json:"content,omitempty" bson:"content,omitempty"`
}