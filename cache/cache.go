package cache

import (
	"fmt"
	"time"
)

// NewCache создает новый объект кэша
//
// max_size - максимальный размер кэша
func NewCache(max_size int) ICache {
	return &Cache{
		data:     make(map[interface{}]CustomValue),
		max_size: max_size,
	}
}

type ICache interface {
	// Len возвращает количество элементов в кэше
	Len() int

	// Cap возвращает максимальный размер кэша
	Cap() int
	// Add добавляет элемент в кэш
	// Если кэш переполнен, удаляет самый старый элемент
	//
	// key - ключ элемента
	// value - значение элемента
	Add(key, value interface{})

	// AddWithTTL добавляет элемент в кэш и задает время жизни элемента
	//
	// key - ключ элемента
	// value - значение элемента
	// ttl - время жизни элемента
	AddWithTTL(key, value interface{}, ttl time.Duration)

	// Clear очищает кэш
	Clear()

	// Get возвращает значение элемента из кэша
	//
	// key - ключ элемента
	//
	// Если элемента нет в кэше, возвращает ok = false
	Get(key interface{}) (interface{}, bool)

	// Remove удаляет элемент из кэша
	//
	// key - ключ элемента
	//
	// Если элемента нет в кэше, возвращает ok = false
	Remove(key interface{}) bool
}

// Cache кэш
//
// data - словарь элементов
// max_size - максимальный размер кэша
type Cache struct {
	data     map[interface{}]CustomValue
	max_size int
}

// Add добавляет элемент в кэш
// Если кэш переполнен, удаляет самый старый элемент
//
// key - ключ элемента
// value - значение элемента
func (c *Cache) Add(key, value interface{}) {
	if len(c.data) == c.max_size {
		fmt.Println("Кэш переполнен, удаляем самый старый элемент")
		var delete_key interface{}
		for k := range c.data {
			delete_key = k
			break
		}
		for k, v := range c.data {
			if v.last_get_time < c.data[delete_key].last_get_time {
				delete_key = k
			}
		}
		c.Remove(delete_key)
	}
	c.data[key] = CustomValue{
		value:         value,
		last_get_time: time.Now().Unix(),
	}
}

// AddWithTTL добавляет элемент в кэш и задает время жизни элемента
//
// key - ключ элемента
// value - значение элемента
// ttl - время жизни элемента
func (c *Cache) AddWithTTL(key, value interface{}, ttl time.Duration) {
	c.Add(key, value)
	c.data[key] = CustomValue{
		value:         value,
		last_get_time: time.Now().Add(ttl).Unix(),
		timer: time.AfterFunc(ttl, func() {
			c.Remove(key)
		}),
	}
}

// Get возвращает значение элемента из кэша
//
// key - ключ элемента
//
// Если элемента нет в кэше, возвращает ok = false
func (c *Cache) Get(key interface{}) (value interface{}, ok bool) {
	if _, ok = c.data[key]; !ok {
		return nil, false
	}
	c.data[key] = CustomValue{
		value:         c.data[key].value,
		last_get_time: time.Now().Unix(),
	}
	return c.data[key].value, ok
}

// Remove удаляет элемент из кэша
//
// key - ключ элемента
//
// Если элемента нет в кэше, возвращает ok = false
func (c *Cache) Remove(key interface{}) (ok bool) {
	if _, ok = c.data[key]; !ok {
		return false
	}
	if v := c.data[key]; v.timer != nil {
		v.timer.Stop()
	}
	delete(c.data, key)
	return true
}

// Clear очищает кэш
func (c *Cache) Clear() {
	for _, v := range c.data {
		if v.timer != nil {
			v.timer.Stop()
		}
	}
	c.data = make(map[interface{}]CustomValue)
}

// Cap возвращает максимальный размер кэша
func (c *Cache) Cap() int {
	return c.max_size
}

// Len возвращает количество элементов в кэше
func (c *Cache) Len() int {
	return len(c.data)
}

// CustomValue структура для хранения значения элемента в кэше
//
// value - значение элемента
// last_get_time - время последнего получения значения
// timer - таймер жизни элемента, если он был добавлен
type CustomValue struct {
	value         interface{}
	last_get_time int64
	timer         *time.Timer
}
