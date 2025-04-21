package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Leaf struct {
	key,
	value int
	parent,
	left,
	right *Leaf
}

type OrderedMap struct {
	root *Leaf
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	m.size++
	if m.root == nil {
		m.root = &Leaf{
			key:   key,
			value: value,
		}

		return
	}

	leaf := m.root.findTerminalLeaf(key)
	if leaf.key > key {
		leaf.left = &Leaf{key: key, value: value, parent: leaf}
	} else {
		leaf.right = &Leaf{key: key, value: value, parent: leaf}
	}
}

func (m *OrderedMap) Erase(key int) {
	if m.root == nil {
		return
	}

	if !m.Contains(key) {
		return
	}
	m.size--

	leaf := m.root.findByKey(key)

	p := leaf.left

	if leaf.right != nil {
		p = leaf.right
	}

	if leaf.left != nil && leaf.right != nil {
		leaf.left.parent = leaf.right
		leaf.right.left = leaf.left
	}

	if leaf.parent == nil {
		m.root = p
		return
	}

	if leaf.parent.key > key {
		leaf.parent.left = p
	} else {
		leaf.parent.right = p
	}

}

func (m *OrderedMap) Contains(key int) bool {
	if m.root == nil {
		return false
	}

	leaf := m.root.findByKey(key)
	if leaf == nil {
		return false
	}
	return true
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	if m.root == nil {
		return
	}

	m.root.foreach(action)
}

func (l *Leaf) foreach(action func(int, int)) {
	if l.left != nil {
		l.left.foreach(action)
	}

	action(l.key, l.value)

	if l.right != nil {
		l.right.foreach(action)
	}
}

func (l *Leaf) findTerminalLeaf(key int) *Leaf {
	if key < l.key {

		if l.left == nil {
			return l
		}

		return l.left.findTerminalLeaf(key)
	} else {

		if l.right == nil {
			return l
		}

		return l.right.findTerminalLeaf(key)
	}
}

func (l *Leaf) findByKey(key int) *Leaf {
	if l.key == key {
		return l
	}

	if key < l.key && l.left != nil {
		return l.left.findByKey(key)
	} else if l.right != nil {
		return l.right.findByKey(key)
	}

	return nil
}

func TestOrderedMap(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	fmt.Println(keys)
	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(10)
	assert.False(t, data.Contains(10))
	assert.Equal(t, 3, data.Size())

	keys = nil
	expectedKeys = []int{4, 5, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	fmt.Println(keys)
	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
