package redis

import (
	"context"
	"fmt"
	"testing"
)

func Test_Alone(t *testing.T) {
	var (
		err  error
		resp string
	)

	cfg := &RedisConfig{
		RedisType: "alone",
		Addrs: []string{
			"10.4.7.71:6379",
		},
		Password: "admin123",
		PoolSize: 100,
	}
	client := NewRedisClient(cfg).Conn
	defer client.Close()

	// call commands on it
	if _, err = client.Set(context.TODO(), "some-key", "2222", 0).Result(); err != nil {
		t.Errorf("SET failed: %v", err)
	}

	if resp, err = client.Get(context.TODO(), "some-key").Result(); err != nil {
		t.Errorf("GET failed: %v", err)
	}

	fmt.Println(resp)
}

func Test_Cluster(t *testing.T) {
	var (
		err  error
		resp string
	)

	cfg := &RedisConfig{
		RedisType: "cluster",
		Addrs: []string{
			"10.4.7.71:6371",
			"10.4.7.71:6372",
			"10.4.7.71:6373",
			"10.4.7.71:6374",
			"10.4.7.71:6375",
			"10.4.7.71:6376",
		},
		Password: "admin123",
		PoolSize: 100,
	}
	client := NewRedisClient(cfg).Conn
	defer client.Close()

	// call commands on it
	if _, err = client.Set(context.TODO(), "some-key", "2222", 0).Result(); err != nil {
		t.Errorf("SET failed: %v", err)
	}

	if resp, err = client.Get(context.TODO(), "some-key").Result(); err != nil {
		t.Errorf("GET failed: %v", err)
	}

	fmt.Println(resp)
}

func Test_Sentinel(t *testing.T) {
	var (
		err  error
		resp string
	)

	cfg := &RedisConfig{
		RedisType: "sentinel",
		Addrs: []string{
			"10.4.7.71:26379",
			"10.4.7.71:26380",
			"10.4.7.71:26381",
		},
		Password:   "admin123",
		PoolSize:   100,
		MasterName: "mymaster",
	}
	client := NewRedisClient(cfg).Conn
	defer client.Close()

	// call commands on it
	if _, err = client.Set(context.TODO(), "some-key", "2222", 0).Result(); err != nil {
		t.Errorf("SET failed: %v", err)
	}

	if resp, err = client.Get(context.TODO(), "some-key").Result(); err != nil {
		t.Errorf("GET failed: %v", err)
	}

	fmt.Println(resp)
}

// 分布式锁
func TestRedisRedsync(t *testing.T) {
	cfg := &RedisConfig{
		RedisType: "cluster",
		Addrs: []string{
			"10.4.7.71:6371",
			"10.4.7.71:6372",
			"10.4.7.71:6373",
			"10.4.7.71:6374",
			"10.4.7.71:6375",
			"10.4.7.71:6376",
		},
		Password: "admin123",
		PoolSize: 100,
	}

	client := NewRedisClient(cfg).Conn
	defer client.Close()

	redisLockClient := NewRedisLock(client, "my-mutex")

	var (
		boolFlag bool
		err      error
	)

	// 增加锁
	if boolFlag, err = redisLockClient.Acquire(); err != nil {
		panic(err)
	}
	fmt.Println("增加锁resp：", boolFlag)

	// 释放锁
	if boolFlag, err = redisLockClient.Release(); err != nil {
		panic(err)
	}
	fmt.Println("释放锁resp：", boolFlag)
}
