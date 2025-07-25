package container

import (
	"context"
	"haircompany-shop-rest/config"
	"haircompany-shop-rest/internal/services"
	"haircompany-shop-rest/pkg/database"
	"sync"
)

type Container struct {
	DB              *database.DB
	JWTService      services.JWTService
	FileService     services.FileSystemService
	PasswordService services.PasswordService
	RedisService    services.RedisService
	Ctx             context.Context
	Wg              *sync.WaitGroup
}

func NewContainer(cfg *config.Config, ctx context.Context, wg *sync.WaitGroup) *Container {
	db := database.NewDB(cfg)
	jwtSvc := services.NewJWTService(cfg.DashboardSecret, cfg.ClientSecret)
	fileSvc := services.NewFileSystemService()
	passwordSvc := services.NewPasswordService()
	redisSvc := services.NewRedisService(ctx, cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)

	return &Container{
		DB:              db,
		JWTService:      jwtSvc,
		FileService:     fileSvc,
		PasswordService: passwordSvc,
		RedisService:    redisSvc,
		Ctx:             ctx,
		Wg:              wg,
	}
}
