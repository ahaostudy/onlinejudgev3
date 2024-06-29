package cache

import (
	"context"
	ktrdb "github.com/ahaostudy/kitextool/option/redis"
	"github.com/google/uuid"
)

type FileCache struct {
	ctx context.Context
}

func NewFileCache(ctx context.Context) *FileCache {
	return &FileCache{ctx: ctx}
}

func (c *FileCache) Store(path string) (string, error) {
	fileId := uuid.New().String()
	err := ktrdb.RDB().Set(c.ctx, fileId, path, 0).Err()
	if err != nil {
		return "", err
	}
	return fileId, nil
}

func (c *FileCache) Load(fileId string) (string, error) {
	path, err := ktrdb.RDB().Get(c.ctx, fileId).Result()
	if err != nil {
		return "", err
	}
	return path, nil
}

func (c *FileCache) Delete(fileId string) error {
	return ktrdb.RDB().Del(c.ctx, fileId).Err()
}
