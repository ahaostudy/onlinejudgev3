package cache

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	ktrdb "github.com/ahaostudy/kitextool/option/server/redis"
)

const (
	FileKey = "file"
)

type FileCache struct {
	ctx context.Context
}

func NewFileCache(ctx context.Context) *FileCache {
	return &FileCache{ctx: ctx}
}

func (c *FileCache) Set(path string) (string, error) {
	fileId := uuid.New().String()
	key := fmt.Sprintf("%s:%s", FileKey, fileId)
	err := ktrdb.RDB().Set(c.ctx, key, path, 0).Err()
	if err != nil {
		return "", err
	}
	return fileId, nil
}

func (c *FileCache) Get(fileId string) (string, error) {
	key := fmt.Sprintf("%s:%s", FileKey, fileId)
	return ktrdb.RDB().Get(c.ctx, key).Result()
}

func (c *FileCache) Del(fileId string) error {
	return ktrdb.RDB().Del(c.ctx, fileId).Err()
}
