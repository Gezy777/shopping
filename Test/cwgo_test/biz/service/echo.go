package service

import (
	"context"
	api "cwgo_test/kitex_gen/api"
	"fmt"

	"github.com/bytedance/gopkg/cloud/metainfo"
)

type EchoService struct {
	ctx context.Context
}

// NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *api.Request) (resp *api.Response, err error) {
	// Finish your business logic.
	clientName, ok := metainfo.GetPersistentValue(s.ctx, "CLIENT_NAME")
	fmt.Println(clientName, ok)
	return &api.Response{Message: req.Message}, nil
}
