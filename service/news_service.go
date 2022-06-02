package service

import (
	"context"
	"go-news/entity"
	"go-news/repository"
	"go-news/web/req"
	"go-news/web/res"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewsService interface {
	FindAll(ctx context.Context) []res.NewsResponse
	Create(ctx context.Context, request req.NewsCreateRequest) res.NewsResponse 
	FindById(ctx context.Context, newsId primitive.ObjectID) res.NewsResponse
	Update(ctx context.Context, request req.NewsUpdateRequest) res.NewsResponse
	Delete(ctx context.Context, newsId primitive.ObjectID)
}

type NewsServiceImpl struct {
	NewsRepository repository.NewsRepository
}

func NewNewsService(newsRepository repository.NewsRepository) NewsService {
	return &NewsServiceImpl{
		NewsRepository: newsRepository,
	}
}

func (newsService *NewsServiceImpl) FindAll(ctx context.Context) []res.NewsResponse {
	lotOfNews := newsService.NewsRepository.FindAll(ctx)
	var newsResponses []res.NewsResponse
	for _, news := range lotOfNews {
		newsResponse := res.NewsResponse{
			Id: news.Id,
			Title: news.Title,
			Content: news.Content,
		}
		newsResponses = append(newsResponses, newsResponse)
	}
	return newsResponses
}

func (newsService *NewsServiceImpl) Create(ctx context.Context, request req.NewsCreateRequest) res.NewsResponse {
	news := entity.News{
		Title: request.Title,
		Content: request.Content,
	}
	news = newsService.NewsRepository.Insert(ctx, news)
	newsResponse := res.NewsResponse{
		Id: news.Id,
		Title: news.Title,
		Content: news.Content,
	}
	return newsResponse
}

func (newsService *NewsServiceImpl) FindById(ctx context.Context, newsId primitive.ObjectID) res.NewsResponse {
	news := newsService.NewsRepository.FindById(ctx, newsId)
	newsResponse := res.NewsResponse{
		Id: news.Id,
		Title: news.Title,
		Content: news.Content,
	}
	return newsResponse
}

func (newsService *NewsServiceImpl) Update(ctx context.Context, request req.NewsUpdateRequest) res.NewsResponse {
	updatedRequest := newsService.NewsRepository.FindById(ctx, request.Id)
	
	updatedRequest.Title = request.Title
	updatedRequest.Content = request.Content
	
	updatedNews := newsService.NewsRepository.Update(ctx, updatedRequest)
	newsResponse := res.NewsResponse{
		Id: updatedNews.Id,
		Title: updatedNews.Title,
		Content: updatedNews.Content,
	}
	return newsResponse
}

func (newsService *NewsServiceImpl) Delete(ctx context.Context, newsId primitive.ObjectID) {
	newsToDelete := newsService.NewsRepository.FindById(ctx, newsId)
	newsService.NewsRepository.Delete(ctx, newsToDelete)
}