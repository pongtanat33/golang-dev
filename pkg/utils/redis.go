package utils

import (
	"bufferbox_backend_go/logs"
	"bufferbox_backend_go/pkg/redis"
	"context"
	"strings"
)

const prefix = "gems-bufferbox"

func RemoveAllDataFromRedisByCompany(ctx context.Context, reqID string) (int, error) {
	lowerReqID := strings.ToLower(reqID)
	registryKey := prefix + ":registry:company:" + lowerReqID

	// 1. ดึงรายการ key ที่อยู่ใน registry set
	keys, err := redis.Rdb.SMembers(ctx, registryKey).Result()
	if err != nil {
		return 0, err
	}

	if len(keys) == 0 {
		logs.Info("No keys found in registry: " + registryKey)
		return 0, nil
	}

	// 2. ลบทุก key ที่อยู่ใน set
	var deleted int64
	for _, key := range keys {
		n, err := redis.Rdb.Del(ctx, key).Result()
		if err != nil {
			return int(deleted), err
		}
		deleted += n
	}

	// 3. ลบ registry set เองด้วย
	_, err = redis.Rdb.Del(ctx, registryKey).Result()
	if err != nil {
		return int(deleted), err
	}
	logs.Info("Deleted registry key: " + registryKey)

	return int(deleted), nil
}
