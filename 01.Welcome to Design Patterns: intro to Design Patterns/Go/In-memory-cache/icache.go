package main

type ICache interface {
	evict(c *cache)
}