package cache

import (
	"b5gocmf/common/daos/system"
	"b5gocmf/utils/core"
	"github.com/go-redis/redis"
)

type ConfigCache struct {
	key string  //redis存储的键
	redis *redis.Client //redis连接对象
}

func NewConfigCache() *ConfigCache {
	return &ConfigCache{
		key: "b5_config_list",
		redis: core.G_Redis.Conn(),
	}
}

func (c *ConfigCache) GetValue(key string) (result string)  {
	result = ""
	if key == "" {
		return
	}

	if res,err:= c.redis.HGet(c.key,key).Result();err!=nil {
		val := system.NewConfigDao().GetInfoByType(key)
		if val == nil {
			return
		}
		result = val.Value
		c.redis.HSet(c.key,key,result)
	}else{
		result = res
	}
	return
}

func (c *ConfigCache) Flush()  {
	core.G_Redis.Conn().Del(c.key)
}

