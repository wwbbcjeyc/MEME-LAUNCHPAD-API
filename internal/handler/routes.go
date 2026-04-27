package handler

import (
	"net/http"

	user "meme-launchpad-api/internal/handler/user"
	"meme-launchpad-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {

	// ==================== 无需认证的接口 ====================

	// User 相关 - 无需认证
	server.AddRoutes(
		[]rest.Route{
			{
				// 获取签名消息
				Method:  http.MethodGet,
				Path:    "/user/sign-msg",
				Handler: user.GetSignMsgHandler(serverCtx),
			},
			{
				// 钱包登录
				Method:  http.MethodPost,
				Path:    "/user/wallet-login",
				Handler: user.WalletLoginHandler(serverCtx),
			},
			{
				// 刷新令牌
				Method:  http.MethodPost,
				Path:    "/user/refresh-token",
				Handler: user.RefreshTokenHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
