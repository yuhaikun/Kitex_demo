// Code generated by Kitex v0.7.0. DO NOT EDIT.

package foundation

import (
	kitex_gen "NewTikTok/kitex_gen"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Feed(ctx context.Context, Req *kitex_gen.DouyinFeedRequest, callOptions ...callopt.Option) (r *kitex_gen.DouyinFeedResponse, err error)
	PublishAction(ctx context.Context, Req *kitex_gen.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *kitex_gen.DouyinPublishActionResponse, err error)
	GetPublishList(ctx context.Context, Req *kitex_gen.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *kitex_gen.DouyinPublishActionResponse, err error)
	GetUserInfo(ctx context.Context, Req *kitex_gen.DouyinUserRequest, callOptions ...callopt.Option) (r *kitex_gen.DouyinUserResponse, err error)
	UserLogin(ctx context.Context, Req *kitex_gen.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *kitex_gen.DouyinUserLoginResponse, err error)
	UserRegister(ctx context.Context, Req *kitex_gen.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *kitex_gen.DouyinUserRegisterResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kFoundationClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFoundationClient struct {
	*kClient
}

func (p *kFoundationClient) Feed(ctx context.Context, Req *kitex_gen.DouyinFeedRequest, callOptions ...callopt.Option) (r *kitex_gen.DouyinFeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Feed(ctx, Req)
}

func (p *kFoundationClient) PublishAction(ctx context.Context, Req *kitex_gen.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *kitex_gen.DouyinPublishActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishAction(ctx, Req)
}

func (p *kFoundationClient) GetPublishList(ctx context.Context, Req *kitex_gen.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *kitex_gen.DouyinPublishActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublishList(ctx, Req)
}

func (p *kFoundationClient) GetUserInfo(ctx context.Context, Req *kitex_gen.DouyinUserRequest, callOptions ...callopt.Option) (r *kitex_gen.DouyinUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, Req)
}

func (p *kFoundationClient) UserLogin(ctx context.Context, Req *kitex_gen.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *kitex_gen.DouyinUserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserLogin(ctx, Req)
}

func (p *kFoundationClient) UserRegister(ctx context.Context, Req *kitex_gen.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *kitex_gen.DouyinUserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserRegister(ctx, Req)
}
