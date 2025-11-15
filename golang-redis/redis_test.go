package golang_redis

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var clientRedis = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "d4v1d4nw4r",
	DB:       0,
})

func TestClientRedis(t *testing.T) {
	assert.NotNil(t, clientRedis)

	err := clientRedis.Close()
	assert.Nil(t, err)
}

var ctx = context.Background()

func TestPing(t *testing.T) {
	result, err := clientRedis.Ping(ctx).Result()
	assert.Nil(t, err)
	assert.Equal(t, "PONG", result)
}

func TestString(t *testing.T) {
	clientRedis.SetEx(ctx, "name", "david", time.Second*2)
	result, err := clientRedis.Get(ctx, "name").Result()
	assert.Nil(t, err)
	assert.Equal(t, "david", result)

	time.Sleep(time.Second * 3)
	result, err = clientRedis.Get(ctx, "name").Result()
	assert.NotNil(t, err)
	assert.Equal(t, redis.Nil, err)
}

func TestList(t *testing.T) {
	clientRedis.RPush(ctx, "names", "Eko")
	clientRedis.RPush(ctx, "names", "Kurniawan")
	clientRedis.RPush(ctx, "names", "Khannedy")

	assert.Equal(t, "Eko", clientRedis.LPop(ctx, "names").Val())
	assert.Equal(t, "Kurniawan", clientRedis.LPop(ctx, "names").Val())
	assert.Equal(t, "Khannedy", clientRedis.LPop(ctx, "names").Val())

	clientRedis.Del(ctx, "names")
}

func TestSet(t *testing.T) {
	clientRedis.SAdd(ctx, "students", "Eko")
	clientRedis.SAdd(ctx, "students", "Eko")
	clientRedis.SAdd(ctx, "students", "Kurniawan")
	clientRedis.SAdd(ctx, "students", "Kurniawan")
	clientRedis.SAdd(ctx, "students", "Khannedy")
	clientRedis.SAdd(ctx, "students", "Khannedy")

	assert.Equal(t, int64(3), clientRedis.SCard(ctx, "students").Val())
	assert.Equal(t, []string{"Eko", "Kurniawan", "Khannedy"}, clientRedis.SMembers(ctx, "students").Val())
}

func TestSortedSet(t *testing.T) {
	clientRedis.ZAdd(ctx, "scores", redis.Z{Score: 100, Member: "Eko"})
	clientRedis.ZAdd(ctx, "scores", redis.Z{Score: 85, Member: "Budi"})
	clientRedis.ZAdd(ctx, "scores", redis.Z{Score: 95, Member: "Joko"})

	assert.Equal(t, []string{"Budi", "Joko", "Eko"}, clientRedis.ZRange(ctx, "scores", 0, -1).Val())

	assert.Equal(t, "Eko", clientRedis.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Joko", clientRedis.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Budi", clientRedis.ZPopMax(ctx, "scores").Val()[0].Member)
}

func TestHash(t *testing.T) {
	clientRedis.HSet(ctx, "user:1", "id", "1")
	clientRedis.HSet(ctx, "user:1", "name", "Eko")
	clientRedis.HSet(ctx, "user:1", "email", "eko@example.com")

	user := clientRedis.HGetAll(ctx, "user:1").Val()

	assert.Equal(t, "1", user["id"])
	assert.Equal(t, "Eko", user["name"])
	assert.Equal(t, "eko@example.com", user["email"])

	clientRedis.Del(ctx, "user:1")
}

func TestHyperLogLog(t *testing.T) {
	clientRedis.PFAdd(ctx, "visitors", "eko", "kurniawan", "khannedy")
	clientRedis.PFAdd(ctx, "visitors", "eko", "budi", "joko")
	clientRedis.PFAdd(ctx, "visitors", "rully", "budi", "joko")

	total := clientRedis.PFCount(ctx, "visitors").Val()
	assert.Equal(t, int64(6), total)
}

func TestPipeline(t *testing.T) {
	_, err := clientRedis.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.SetEx(ctx, "name", "Eko", 5*time.Second)
		pipeliner.SetEx(ctx, "address", "Indonesia", 5*time.Second)
		return nil
	})
	assert.Nil(t, err)

	assert.Equal(t, "Eko", clientRedis.Get(ctx, "name").Val())
	assert.Equal(t, "Indonesia", clientRedis.Get(ctx, "address").Val())
}

func TestTransaction(t *testing.T) {
	_, err := clientRedis.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.SetEx(ctx, "name", "Joko", 5*time.Second)
		pipeliner.SetEx(ctx, "address", "Cirebon", 5*time.Second)
		return nil
	})
	assert.Nil(t, err)

	assert.Equal(t, "Joko", clientRedis.Get(ctx, "name").Val())
	assert.Equal(t, "Cirebon", clientRedis.Get(ctx, "address").Val())
}

func TestPublishStream(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := clientRedis.XAdd(ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]interface{}{
				"name":    "Eko",
				"address": "Indonesia",
			},
		}).Err()
		assert.Nil(t, err)
	}
}

func TestCreateConsumerGroup(t *testing.T) {
	clientRedis.XGroupCreate(ctx, "members", "group-1", "0")
	clientRedis.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-1")
	clientRedis.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-2")
}

func TestConsumeStream(t *testing.T) {
	streams := clientRedis.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    "group-1",
		Consumer: "consumer-1",
		Streams:  []string{"members", ">"},
		Count:    2,
		Block:    5 * time.Second,
	}).Val()

	for _, stream := range streams {
		for _, message := range stream.Messages {
			fmt.Println(message.ID)
			fmt.Println(message.Values)
		}
	}
}

func TestSubscribePubSub(t *testing.T) {
	subscriber := clientRedis.Subscribe(ctx, "channel-1")
	defer subscriber.Close()
	for i := 0; i < 10; i++ {
		message, err := subscriber.ReceiveMessage(ctx)
		assert.Nil(t, err)
		fmt.Println(message.Payload)
	}
}

func TestPublishPubSub(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := clientRedis.Publish(ctx, "channel-1", "Hello "+strconv.Itoa(i)).Err()
		assert.Nil(t, err)
	}
}
