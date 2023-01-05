package logic

import (
	"context"

	"live/internal/svc"
	"live/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LiveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveLogic {
	return &LiveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LiveLogic) Live(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	roomInfo, err := l.svcCtx.HuYa.GetLiveUrl(l.ctx, req.RoomId)
	if err != nil {
		return nil, err
	}
	logx.Infof("liveUrl:", roomInfo)
	resp = &types.Response{
		Urls:    roomInfo.Urls,
		Name:    roomInfo.Name,
		LiveUrl: roomInfo.LiveUrl,
	}
	return
}
