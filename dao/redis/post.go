package redis

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
	"web_app/models"
)

func getIDsFromKey(key string, page int64, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	// ZREVRANGE 按分数从大到小的顺序查询指定数量的元素
	return rdb.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取id
	// 1. 根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 2. 确定查询的索引起始点
	return getIDsFromKey(key, p.Page, p.Size)
}

func GetCommunityPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 获取排序key
	orderkey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderkey = getRedisKey(KeyPostScoreZSet)
	}
	// 使用 zinterstore 把分区的帖子set与帖子分数的 zset 生成一个新的zset
	// 针对新的zset 按之前的逻辑取数据

	//社区的key
	cKey := getRedisKey(KeyCommunitySetPF) + strconv.Itoa(int(p.CommunityID))

	key := orderkey + strconv.Itoa(int(p.CommunityID))
	//利用缓存key减少zinterstore执行的次数
	if rdb.Exists(key).Val() < 1 {
		//创建set
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, orderkey, cKey)
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	// 存在则直接查询缓存
	return getIDsFromKey(key, p.Page, p.Size)
}
