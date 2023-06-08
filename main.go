package main

import (
	"fmt"
	"time"

	"icache/cache"
)

func main() {
	c := cache.NewCache(5)
	c.AddWithTTL("one", 1, time.Second)
	<-time.After(time.Second)
	c.Add("two", 2)
	<-time.After(time.Second)
	c.Add("three", 3)
	<-time.After(time.Second)
	c.Add("four", 4)
	<-time.After(time.Second)
	c.Add("five", 5)
	<-time.After(time.Second)
	fmt.Println(c)
	fmt.Println(c.Get("one"))
	c.Add("six", 6)
	fmt.Println(c)
}
