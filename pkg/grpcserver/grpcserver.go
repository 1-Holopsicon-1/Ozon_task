package grpcserver

import (
	"context"
	api "github.com/1-Holopsicon-1/Ozon_task/pkg/api"
)

type GRPCSever struct {}

func (s *GRPCSever) Create(ctx context.Context, req *api.OldUrl) (*api.CutUrl, error) {
	return &api.CutUrl{NewUrl: req.GetOldUrl()}, nil
}