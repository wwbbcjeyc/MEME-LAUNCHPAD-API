// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret  string   // 签名密钥（最重要！）
		AccessExpire  int64    // Access Token 过期时间（秒）
		RefreshExpire int64    // Refresh Token 过期时间（秒）
	}

	Database struct {
		Host         string
		Port         int
		User         string
		Password     string
		Name         string
		SSLMode      string
		MaxOpenConns int
		MaxIdleConns int
	}

	Redis struct {
		Host string
		Pass string
		DB   int
	}
}
