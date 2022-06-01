package todo

import (
	"context"
	"github.com/todoApp/pkg/storage"
	"github.com/todoApp/proto"
	"time"
)

type Api interface {
	AddItem(ctx context.Context, request *proto.AddItemRequest) (*proto.AddItemResponse, error)
	GetItem(ctx context.Context, request *proto.GetSingleItemRequest) (*proto.GetSingleItemResponse, error)
	GetAllItems(ctx context.Context, request *proto.GetAllItemsRequest) (*proto.GetAllItemsResponse, error)
	UpdateItemDescription(ctx context.Context, request *proto.UpdateItemDescriptionRequest) (*proto.UpdateItemDescriptionResponse, error)
	UpdateItemCompletion(ctx context.Context, request *proto.UpdateItemCompletionRequest) (*proto.UpdateItemCompletionResponse, error)
	DeleteItem(ctx context.Context, request *proto.DeleteItemRequest) (*proto.DeleteItemResponse, error)
}

type TodolistService struct {
	db             storage.Storage
	requestTimeout time.Duration
}

func NewTodoService(db storage.Storage, requestTimeout int64) Api {
	return &TodolistService{
		db:             db,
		requestTimeout: time.Duration(requestTimeout) * time.Second,
	}
}

func (t *TodolistService) AddItem(ctx context.Context, request *proto.AddItemRequest) (*proto.AddItemResponse, error) {
	item, err := t.db.AddItem(request.GetItem().GetDescription())
	if err != nil {
		return nil, err
	}

	response := &proto.AddItemResponse{
		Item: &proto.Item{
			Id:          uint32(item.ID),
			Description: item.Description,
			Completed:   item.Completed,
		}}

	return response, nil
}

func (t *TodolistService) GetItem(ctx context.Context, request *proto.GetSingleItemRequest) (*proto.GetSingleItemResponse, error) {
	item, err := t.db.GetSingleItem(request.GetId())
	if err != nil {
		return nil, err
	}

	response := &proto.GetSingleItemResponse{
		Item: &proto.Item{
			Id:          uint32(item.ID),
			Description: item.Description,
			Completed:   item.Completed,
		}}

	return response, nil
}

func (t *TodolistService) GetAllItems(ctx context.Context, request *proto.GetAllItemsRequest) (*proto.GetAllItemsResponse, error) {
	items, err := t.db.GetItems()
	if err != nil {
		return nil, err
	}

	var responseItems []*proto.Item
	for _, item := range items {
		responseItems = append(responseItems, &proto.Item{
			Id:          uint32(item.ID),
			Description: item.Description,
			Completed:   item.Completed,
		})
	}

	response := &proto.GetAllItemsResponse{Items: responseItems}
	return response, nil
}

func (t *TodolistService) UpdateItemDescription(ctx context.Context, request *proto.UpdateItemDescriptionRequest) (*proto.UpdateItemDescriptionResponse, error) {
	item, err := t.db.UpdateItemDescription(request.GetId(), request.GetDescription())
	if err != nil {
		return nil, err
	}

	response := &proto.UpdateItemDescriptionResponse{
		Item: &proto.Item{
			Id:          uint32(item.ID),
			Description: item.Description,
			Completed:   item.Completed,
		},
	}

	return response, nil
}

func (t *TodolistService) UpdateItemCompletion(ctx context.Context, request *proto.UpdateItemCompletionRequest) (*proto.UpdateItemCompletionResponse, error) {
	item, err := t.db.UpdateItemCompletion(request.GetId(), request.GetCompleted())
	if err != nil {
		return nil, err
	}

	response := &proto.UpdateItemCompletionResponse{
		Item: &proto.Item{
			Id:          uint32(item.ID),
			Description: item.Description,
			Completed:   item.Completed,
		},
	}

	return response, nil
}

func (t *TodolistService) DeleteItem(ctx context.Context, request *proto.DeleteItemRequest) (*proto.DeleteItemResponse, error) {
	item, err := t.db.DeleteItem(request.GetId())
	if err != nil {
		return nil, err
	}

	response := &proto.DeleteItemResponse{
		Item: &proto.Item{
			Id:          uint32(item.ID),
			Description: item.Description,
			Completed:   item.Completed,
		},
	}

	return response, nil
}
