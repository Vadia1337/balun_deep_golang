package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values []int
	nextPushIdx,
	nextPopIdx int
}

func NewCircularQueue(size int) *CircularQueue {
	if size <= 0 {
		return nil
	}

	values := make([]int, size)
	for i := range values {
		values[i] = -1
	}

	return &CircularQueue{
		values: values,
	}
}

func (q *CircularQueue) Push(value int) bool {
	if q.values[q.nextPushIdx] != -1 {
		return false
	}
	q.values[q.nextPushIdx] = value

	if q.nextPushIdx+1 == len(q.values) {
		q.nextPushIdx = 0
	} else {
		q.nextPushIdx++
	}

	return true
}

func (q *CircularQueue) Pop() bool {
	if q.values[q.nextPopIdx] == -1 {
		return false
	}
	q.values[q.nextPopIdx] = -1

	if q.nextPopIdx+1 == len(q.values) {
		q.nextPopIdx = 0
	} else {
		q.nextPopIdx++
	}

	return true
}

func (q *CircularQueue) Front() int {
	return q.values[q.nextPopIdx]
}

func (q *CircularQueue) Back() int {
	if q.nextPushIdx-1 < 0 {
		return q.values[len(q.values)-1]
	} else {
		return q.values[q.nextPushIdx-1]
	}
}

func (q *CircularQueue) Empty() bool {
	if q.nextPushIdx == q.nextPopIdx && q.values[q.nextPushIdx] == -1 {
		return true
	}

	return false
}

func (q *CircularQueue) Full() bool {
	if q.nextPushIdx == q.nextPopIdx && q.values[q.nextPushIdx] != -1 {
		return true
	}

	return false
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
