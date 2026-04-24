// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"meme-launchpad-api/internal/svc"
	"meme-launchpad-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSignMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取签名消息
func NewGetSignMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSignMsgLogic {
	return &GetSignMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSignMsgLogic) GetSignMsg(req *types.GetSignMsgRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
