package main

import (
	kitex_gen "NewTikTok/kitex_gen"
	"context"
)

// AdderImpl implements the last service interface defined in the IDL.
type AdderImpl struct{}

// Add implements the AdderImpl interface.
func (s *AdderImpl) Add(ctx context.Context, req *kitex_gen.AddRequest) (resp *kitex_gen.AddResponse, err error) {
	// TODO: Your code here...
	resp = &kitex_gen.AddResponse{Res: req.X + req.Y}
	return
}
