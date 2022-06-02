package req

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewsUpdateRequest struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id"`
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
}