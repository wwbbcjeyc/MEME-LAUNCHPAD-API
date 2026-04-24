// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"meme-launchpad-api/internal/svc"
	"meme-launchpad-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WalletLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 钱包登录
func NewWalletLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WalletLoginLogic {
	return &WalletLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WalletLoginLogic) WalletLogin(req *types.WalletLoginRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
