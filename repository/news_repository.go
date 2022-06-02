package repository

import (
	"context"
	"go-news/entity"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NewsRepository interface {
	FindAll(ctx context.Context) []entity.News
	Insert(ctx context.Context, news entity.News) entity.News
	FindById(ctx context.Context, newsId primitive.ObjectID) entity.News
	Update(ctx context.Context, news entity.News) entity.News
	Delete(ctx context.Context, news entity.News)
}

type NewsRepositoryImpl struct {
	mongoDB *mongo.Database
}

func NewNewsRepository(DB *mongo.Database) NewsRepository {
	return &NewsRepositoryImpl{
		mongoDB: DB,
	}
}

func (newsRepository *NewsRepositoryImpl) FindAll(ctx context.Context) []entity.News {
	cursor, err := newsRepository.mongoDB.Collection("news").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var lotOfNews []entity.News
	for cursor.Next(ctx) {
		var news entity.News
		err := cursor.Decode(&news)
		if err != nil {
			log.Fatal(err)
		}
		lotOfNews = append(lotOfNews, news)
	}

	return lotOfNews
}

func (newsRepository *NewsRepositoryImpl) Insert(ctx context.Context, news entity.News) entity.News {
	result, err := newsRepository.mongoDB.Collection("news").InsertOne(ctx, news)
	if err != nil {
		log.Fatal(err)
	}
	id := result.InsertedID
	oid := id.(primitive.ObjectID)
	news.Id = oid
	return news
}

func (newsRepository *NewsRepositoryImpl) FindById(ctx context.Context, newsId primitive.ObjectID) entity.News {
	cursor := newsRepository.mongoDB.Collection("news").FindOne(ctx, bson.D{{Key: "_id", Value: newsId}})

	var news entity.News
	
	err := cursor.Decode(&news)
	if err != nil {
		log.Fatal(err)
	}
	return news
}

func (newsRepository *NewsRepositoryImpl) Update(ctx context.Context, news entity.News) entity.News {
	filter := bson.D{{"_id", news.Id}}
	updatedNews := bson.D{{"$set", bson.D{{"title", news.Title}, {"content", news.Content}}}}
	_, err := newsRepository.mongoDB.Collection("news").UpdateOne(ctx, filter, updatedNews)
	if err != nil {
		log.Fatal(err)
	}
	return news
}

func (newsRepository *NewsRepositoryImpl) Delete(ctx context.Context, news entity.News) {
	filter := bson.D{{"_id", news.Id}}
	_, err := newsRepository.mongoDB.Collection("news").DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
}