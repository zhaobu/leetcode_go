package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCircularDeque_InsertFront(t *testing.T) {
	t.Run("测试1", func(t *testing.T) {
		obj := Constructor(3) // 设置容量大小为3
		circularDeque := &obj
		assert.Equal(t, true, circularDeque.InsertLast(1))   // 返回 true
		assert.Equal(t, true, circularDeque.InsertLast(2))   // 返回 true
		assert.Equal(t, true, circularDeque.InsertFront(3))  // 返回 true
		assert.Equal(t, false, circularDeque.InsertFront(4)) // 已经满了，返回 false
		assert.Equal(t, 2, circularDeque.GetRear())          // 返回 2
		assert.Equal(t, true, circularDeque.IsFull())        // 返回 true
		assert.Equal(t, true, circularDeque.DeleteLast())    // 返回 true
		assert.Equal(t, true, circularDeque.InsertFront(4))  // 返回 true
		assert.Equal(t, 4, circularDeque.GetFront())         // 返回 4
	})
}
