package test

import (
	"errors"
	"icache/cache"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	want_len := 5
	c := cache.NewCache(5)
	c.Add("one", 1)
	c.Add("two", 2)
	c.Add(3, "three")
	c.Add(4, 4.44)
	c.Add(5, errors.New("five"))
	c.Add("six", 6)
	if c.Len() != want_len {
		t.Errorf("Кэш должен содержать %d элементов", want_len)
	}
}

func TestGet(t *testing.T) {
	want := 1
	c := cache.NewCache(5)
	c.Add("one", 1)
	c.Add("two", 2)
	c.Add(3, "three")
	test_get, ok := c.Get("one")
	if !ok {
		t.Errorf("Кэш не содержит элемент")
	}
	if test_get != want {
		t.Errorf("Кэш должен содержать %d", want)
	}
}

func TestRemove(t *testing.T) {
	want_len := 4
	c := cache.NewCache(5)
	c.Add("one", 1)
	c.Add("two", 2)
	c.Add(3, "three")
	c.Add(4, 4.44)
	c.Add(5, errors.New("five"))
	c.Remove(5)
	if c.Len() != want_len {
		t.Errorf("Кэш должен содержать %d элементов", want_len)
	}
}

func TestClear(t *testing.T) {
	want_len := 0
	c := cache.NewCache(5)
	c.Add("one", 1)
	c.Add("two", 2)
	c.Add(3, "three")
	c.Add(4, 4.44)
	c.Add(5, errors.New("five"))
	c.Clear()
	if c.Len() != want_len {
		t.Errorf("Кэш должен содержать %d элементов", want_len)
	}
}

func TestAddWithTTL(t *testing.T) {
	want_len := 4
	c := cache.NewCache(5)
	c.AddWithTTL("one", 1, time.Second)
	c.Add("two", 2)
	c.Add(3, "three")
	c.Add(4, 4.44)
	c.Add(5, errors.New("five"))
	<-time.After(time.Second * 2)
	if c.Len() != want_len {
		t.Errorf("Кэш должен содержать %d элементов", want_len)
	}
}
