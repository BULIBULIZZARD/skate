package redis

import (
	"file/skate/config"
	"github.com/garyburd/redigo/redis"
)

type Redis struct {
}

func NewRedis() *Redis {
	return new(Redis)
}

func (r *Redis) getRedis() (redis.Conn, error) {
	redisConfig := config.GetConfig().GetRedisConfig()
	c, err := redis.Dial("tcp", redisConfig.IP+":"+redisConfig.Port)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *Redis) SetValue(key string, value string,outTime string) error {
	c, err := r.getRedis()
	if err != nil {
		return err
	}
	_, err = c.Do("Set", key, value, "EX", "100")
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) GetValue(key string) (string, error) {
	c, err := r.getRedis()
	if err != nil {
		return "", err
	}
	value, err := redis.String(c.Do("Get", key ))
	if err != nil {
		return "", err
	}
	return value, nil
}
