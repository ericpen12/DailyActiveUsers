package main

import (
	"testing"
)

func TestNewRedisClient(t *testing.T) {
	if err := rdb.Ping().Err(); err != nil {
		t.Errorf("test connect redis: %s", err)
	}
}

func TestSign(t *testing.T) {
	if err := Sign(1); err != nil {
		t.Error(err)
	}
}

func TestGetActiveUserCount(t *testing.T) {
	if count, err := GetActiveUserCount(); err != nil {
		t.Error(err)
	}else {
		t.Log(count)
	}
}
