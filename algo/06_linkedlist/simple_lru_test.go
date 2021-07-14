package _6_linkedlist

import (
	"leetcode/utils/zaplog"
	"testing"
)

func init() {
	zaplog.InitLog(&zaplog.Config{
		Logdir:   "./",
		LogName:  "main.log",
		Loglevel: "debug",
		Debug:    true},
	)
}

func TestLRUCache1(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	res := cache.Get(1)
	zaplog.Infof("get(%d) res=%d", 1, res)

	cache.Put(3, 3)

	res = cache.Get(2)
	zaplog.Infof("get(%d) res=%d", 2, res)
	cache.Put(4, 4)

	res = cache.Get(1)
	zaplog.Infof("get(%d) res=%d", 1, res)

	res = cache.Get(3)
	zaplog.Infof("get(%d) res=%d", 3, res)

	res = cache.Get(4)
	zaplog.Infof("get(%d) res=%d", 4, res)
}

func TestLRUCache2(t *testing.T) {
	cache := Constructor(1)
	cache.Put(2, 1)

	res := cache.Get(2)
	zaplog.Infof("get(%d) res=%d", 2, res)

	cache.Put(3, 2)

	res = cache.Get(2)
	zaplog.Infof("get(%d) res=%d", 2, res)
	cache.Put(4, 4)

	res = cache.Get(3)
	zaplog.Infof("get(%d) res=%d", 3, res)

}

/*
["LRUCache","put","get","put","get","get"]
[[1],[2,1],[2],[3,2],[2],[3]]
*/
func TestLRUCache3(t *testing.T) {
	cache := Constructor(1)
	cache.Put(2, 1)

	res := cache.Get(2)
	zaplog.Infof("get(%d) res=%d", 2, res)

	cache.Put(3, 2)

	res = cache.Get(2)
	zaplog.Infof("get(%d) res=%d", 2, res)
	cache.Put(4, 4)

	res = cache.Get(3)
	zaplog.Infof("get(%d) res=%d", 3, res)

}

/*
["LRUCache","put","put","get","put","put","get"]
[[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
*/
func TestLRUCache4(t *testing.T) {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)

	res := cache.Get(2)
	zaplog.Infof("get(%d) res=%d", 2, res)

	cache.Put(1, 1)
	cache.Put(4, 1)

	res = cache.Get(2)
	zaplog.Infof("get(%d) res=%d", 2, res)
}
