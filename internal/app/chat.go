package app

import (
	"context"
	desc "github.com/croked91/chat_api"
	"github.com/golang/protobuf/ptypes/empty"
)

type Chat struct {
	desc.UnimplementedUserAPIServer
}

func (chat *Chat) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	res := &desc.CreateResponse{Id: 1}
	return res, nil
}

func (chat *Chat) Delete(ctx context.Context, request *desc.DeleteRequest) (*empty.Empty, error) {
	return nil, nil
}

func (chat *Chat) SendMessage(ctx context.Context, request *desc.SendMessageRequest) (*empty.Empty, error) {
	return nil, nil
}
