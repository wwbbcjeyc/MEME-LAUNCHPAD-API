# MEME Launchpad API

基于 go-zero 框架的 MEME 代币发射平台后端 API 服务。

# 项目结构
```
MEME-LAUNCHPAD-API/
├── main.go                    # 程序入口
├── go.mod / go.sum            # Go 模块依赖
├── Makefile                   # 构建脚本
├── README.md                  # 项目说明
├── api/                       # API 定义文件（goctl 生成的 API 描述）
│   ├── api.api               # API 接口定义
│   └── types.api             # API 类型定义
├── etc/                       # 配置文件
│   └── api.yaml              # 主配置文件
├── internal/                  # 核心业务代码
│   ├── config/               # 配置定义
│   │   └── config.go
│   ├── handler/              # HTTP 处理器层（接收请求）
│   │   ├── routes.go         # 路由注册
│   │   ├── token/           # Token 相关处理器
│   │   └── user/            # 用户相关处理器
│   ├── logic/                # 业务逻辑层（核心业务处理）
│   │   ├── kline/
│   │   ├── token/
│   │   └── user/
│   ├── middleware/           # 中间件
│   │   └── authmiddleware.go # JWT 认证中间件
│   ├── model/                # 数据模型层（数据库模型）
│   │   ├── activity.go       # 活动模型
│   │   ├── comment.go        # 评论模型
│   │   ├── invite.go         # 邀请模型
│   │   ├── kline.go          # K线数据模型
│   │   ├── token.go          # 代币模型
│   │   ├── trade.go          # 交易模型
│   │   └── user.go           # 用户模型
│   ├── service/              # 外部服务封装
│   │   ├── chain/            # 区块链交互服务
│   │   │   └── events.go
│   │   ├── cos/              # 腾讯云 COS 对象存储
│   │   │   └── cos.go
│   │   ├── crypto/           # 加密相关（CREATE2、Signer）
│   │   │   ├── create2.go
│   │   │   ├── signer.go
│   │   │   └── test_encoding.go
│   │   └── token/            # Token 服务
│   │       └── service.go
│   ├── svc/                   # 服务上下文
│   │   └── servicecontext.go
│   └── types/                # 通用类型定义
│       └── types.go
└── migrations/               # 数据库迁移脚本
    ├── add_available_tokens.sql
    └── init.sql
```


```bash
go mod tidy
```

## 1. 初始化数据库

```bash
make migrate
```

### 2. 运行

```bash
# 开发模式
make dev

# 或编译后运行
make run
```


## Docker 部署

```bash
# 构建镜像
make docker-build

# 运行容器
make docker-run
```