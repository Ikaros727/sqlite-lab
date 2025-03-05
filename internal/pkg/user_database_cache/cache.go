package user_database_cache

import (
	"fmt"
	"github.com/Ikaros727/sqlite-lab/internal/repo"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"time"
)

const (
	// userDBKeyFormat 用户数据库连接Key，例：<userID>/<database>
	userDBKeyFormat = "%d/%s"
)

type UserDatabaseCache struct {
	c *cache.Cache
}

func (u *UserDatabaseCache) Store(userID int64, database string, db *gorm.DB) error {
	return u.c.Add(fmt.Sprintf(userDBKeyFormat, userID, database), db, cache.DefaultExpiration)
}

func (u *UserDatabaseCache) Load(userID int64, database string) (db *gorm.DB, found bool) {
	inst, found := u.c.Get(fmt.Sprintf(userDBKeyFormat, userID, database))
	if !found {
		return nil, false
	}

	return inst.(*gorm.DB), true
}

func (u *UserDatabaseCache) MustLoad(userID int64, database string) (db *gorm.DB, err error) {
	key := fmt.Sprintf(userDBKeyFormat, userID, database)
	inst, found := u.c.Get(key)
	if !found {
		inst, err = repo.NewSQLite(database)
		if err != nil {
			return
		}
	}
	u.c.Set(key, inst, cache.DefaultExpiration)
	db = inst.(*gorm.DB)
	return
}

func NewUserDatabaseCache(defaultExpiration, cleanupInterval time.Duration) *UserDatabaseCache {
	return &UserDatabaseCache{
		c: cache.New(defaultExpiration, cleanupInterval),
	}
}
