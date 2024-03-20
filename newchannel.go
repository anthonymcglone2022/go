package main

import "fmt"

type BasicQueue interface {
	enQueue(val int) string
	deQueue() string
}

type channelAsAQueueObject struct {
	channel chan int
}

func createQueue(size int) BasicQueue {
	return &channelAsAQueueObject{
		channel: make(chan int, size),
	}
}

func (a *channelAsAQueueObject) enQueue(value int) string {
	select {
		case a.channel <- value:
			return fmt.Sprintf("Added to Queue %d", value) 
		default:
			return "Queue Full"
	}
}

func (a *channelAsAQueueObject) deQueue() string {
	select {
		case value := <- a.channel:
			return fmt.Sprintf("Removed from Queue %d", value) 
		default:
			return "Queue Empty"
	}
}

func main() {
	channelBackedQueue := createQueue(3)
	fmt.Println(channelBackedQueue.enQueue(1))
	fmt.Println(channelBackedQueue.enQueue(2))
	fmt.Println(channelBackedQueue.enQueue(3))
	fmt.Println(channelBackedQueue.enQueue(4))
}