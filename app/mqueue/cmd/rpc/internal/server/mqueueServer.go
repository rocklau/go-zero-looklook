// Code generated by goctl. DO NOT EDIT!
// Source: mqueue.proto

package server

import (
	"context"

	"looklook/app/mqueue/cmd/rpc/internal/logic"
	"looklook/app/mqueue/cmd/rpc/internal/svc"
	"looklook/app/mqueue/cmd/rpc/pb"
)

type MqueueServer struct {
	svcCtx *svc.ServiceContext
}

func NewMqueueServer(svcCtx *svc.ServiceContext) *MqueueServer {
	return &MqueueServer{
		svcCtx: svcCtx,
	}
}

// 添加民宿订单延迟关闭到asynq队列
func (s *MqueueServer) AqDeferHomestayOrderClose(ctx context.Context, in *pb.AqDeferHomestayOrderCloseReq) (*pb.AqDeferHomestayOrderCloseResp, error) {
	l := logic.NewAqDeferHomestayOrderCloseLogic(ctx, s.svcCtx)
	return l.AqDeferHomestayOrderClose(in)
}

// 支付流水状态变更发送到kq
func (s *MqueueServer) KqPaymenStatusUpdate(ctx context.Context, in *pb.KqPaymenStatusUpdateReq) (*pb.KqPaymenStatusUpdateResp, error) {
	l := logic.NewKqPaymenStatusUpdateLogic(ctx, s.svcCtx)
	return l.KqPaymenStatusUpdate(in)
}

// 发送微信小程序订阅消息
func (s *MqueueServer) SendWxMiniSubMessage(ctx context.Context, in *pb.SendWxMiniSubMessageReq) (*pb.SendWxMiniSubMessageResp, error) {
	l := logic.NewSendWxMiniSubMessageLogic(ctx, s.svcCtx)
	return l.SendWxMiniSubMessage(in)
}
