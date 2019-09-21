package main

import (
	"github.com/xiazemin/lru/hashlist"
	"fmt"
)

func main(){
	lruCache :=hashlist.Constructor(4)
	lruCache.Put(1,1)
	lruCache.Put(2,2)
	lruCache.Put(3,3)
	lruCache.Put(4,4)
	lruCache.Put(5,5)
	lruCache.Get(1)
	lruCache.Put(6,6)
	lruCache.Put(7,7)
	lruCache.Put(5,5)
	fmt.Println(lruCache.Get(1))


	fmt.Println(lruCache.Get(6))
	fmt.Print(lruCache.HashMap)
}
