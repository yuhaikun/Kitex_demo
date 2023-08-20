package main

import (
	kitex_gen "NewTikTok/kitex_gen"
	"context"
)

// FoundationImpl implements the last service interface defined in the IDL.
type FoundationImpl struct{}

// Feed implements the FoundationImpl interface.
func (s *FoundationImpl) Feed(ctx context.Context, req *kitex_gen.DouyinFeedRequest) (resp *kitex_gen.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishAction implements the FoundationImpl interface.
func (s *FoundationImpl) PublishAction(ctx context.Context, req *kitex_gen.DouyinPublishActionRequest) (resp *kitex_gen.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetPublishList implements the FoundationImpl interface.
func (s *FoundationImpl) GetPublishList(ctx context.Context, req *kitex_gen.DouyinPublishActionRequest) (resp *kitex_gen.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the FoundationImpl interface.
func (s *FoundationImpl) GetUserInfo(ctx context.Context, req *kitex_gen.DouyinUserRequest) (resp *kitex_gen.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}

// UserLogin implements the FoundationImpl interface.
func (s *FoundationImpl) UserLogin(ctx context.Context, req *kitex_gen.DouyinUserLoginRequest) (resp *kitex_gen.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UserRegister implements the FoundationImpl interface.
func (s *FoundationImpl) UserRegister(ctx context.Context, req *kitex_gen.DouyinUserRegisterRequest) (resp *kitex_gen.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}
