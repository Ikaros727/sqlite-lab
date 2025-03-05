package svc

import (
	"github.com/Ikaros727/sqlite-lab/internal/config"
	"github.com/Ikaros727/sqlite-lab/internal/pkg/user_database_cache"
	"time"
)

type ServiceContext struct {
	Config            config.Config
	UserDatabaseCache *user_database_cache.UserDatabaseCache
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		UserDatabaseCache: user_database_cache.NewUserDatabaseCache(5*time.Minute, 10*time.Minute),
	}
}
