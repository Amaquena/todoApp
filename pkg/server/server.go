package server

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/todoApp/pkg/api/todo"
	"github.com/todoApp/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"

	"github.com/todoApp/proto"
)

var log = logrus.WithField("ctx", "server")

type Server struct {
	grpcServer *grpc.Server
	todoApi    todo.Api
}

func NewServer(todoService todo.Api) *Server {
	var opts []grpc.ServerOption

	return &Server{
		grpcServer: grpc.NewServer(opts...),
		todoApi:    todoService,
	}
}

func (s *Server) Serve(conf *config.Server) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", conf.Host, conf.Port))
	if err != nil {
		log.WithFields(logrus.Fields{
			"host": conf.Host,
			"port": conf.Port,
		}).Fatal("Failed to start server")
	}

	proto.RegisterTodoAppAPIServer(s.grpcServer, s)
	reflection.Register(s.grpcServer)

	go func() {
		if err = s.grpcServer.Serve(listener); err != nil {
			log.WithError(err).Fatal("Failed to start server")
		}
	}()
}

func (s *Server) Shutdown() {
	s.grpcServer.GracefulStop()
}

func (s *Server) AddItem(ctx context.Context, request *proto.AddItemRequest) (*proto.AddItemResponse, error) {
	validationErrors := request.ValidateAll()
	if validationErrors != nil {
		log.WithError(validationErrors).Error("AddItem validation failed")
		return nil, status.Error(codes.InvalidArgument, "AddItemRequest validation failed")
	}

	response, err := s.todoApi.AddItem(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}

func (s *Server) GetItem(ctx context.Context, request *proto.GetSingleItemRequest) (*proto.GetSingleItemResponse, error) {
	validationErrors := request.ValidateAll()
	if validationErrors != nil {
		log.WithError(validationErrors).Error("GetItem validation failed")
		return nil, status.Error(codes.InvalidArgument, "GetItemRequest validation failed")
	}

	item, err := s.todoApi.GetItem(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return item, nil
}

func (s *Server) GetAllItems(ctx context.Context, request *proto.GetAllItemsRequest) (*proto.GetAllItemsResponse, error) {
	items, err := s.todoApi.GetAllItems(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return items, nil
}

func (s *Server) UpdateItemDescription(ctx context.Context, request *proto.UpdateItemDescriptionRequest) (*proto.UpdateItemDescriptionResponse, error) {
	validationErrors := request.ValidateAll()
	if validationErrors != nil {
		log.WithError(validationErrors).Error("UpdateItemDescription validation failed")
		return nil, status.Error(codes.InvalidArgument, "UpdateItemDescription validation failed")
	}

	item, err := s.todoApi.UpdateItemDescription(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return item, nil
}

func (s *Server) UpdateItemCompletion(ctx context.Context, request *proto.UpdateItemCompletionRequest) (*proto.UpdateItemCompletionResponse, error) {
	validationErrors := request.ValidateAll()
	if validationErrors != nil {
		log.WithError(validationErrors).Error("UpdateItemCompletion validation failed")
		return nil, status.Error(codes.InvalidArgument, "UpdateItemCompletion validation failed")
	}

	item, err := s.todoApi.UpdateItemCompletion(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return item, nil
}

func (s *Server) DeleteItem(ctx context.Context, request *proto.DeleteItemRequest) (*proto.DeleteItemResponse, error) {
	item, err := s.todoApi.DeleteItem(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return item, nil
}
