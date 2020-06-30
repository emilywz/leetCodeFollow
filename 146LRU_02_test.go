package main

import (
	"testing"
	"fmt"
	)

func TestLRU(t *testing.T) {
	lruCache := Constructor(3)

	lruCache.Put(1, 11)
	lruCache.Put(2, 22)
	lruCache.Put(3, 33)
	lruCache.Put(4, 44)

	aa := lruCache.Get(3)
	fmt.Println("aa:", aa)

	bb := lruCache.Get(4)
	fmt.Println("bb", bb)

	cc := lruCache.Get(1)
	fmt.Println("cc", cc)
}