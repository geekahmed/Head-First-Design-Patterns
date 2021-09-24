package main
import "fmt"
type lfu struct{}

func (l *lru) evict(c *cache){
	fmt.Println("This is the LFU algorithm implementation")
}
