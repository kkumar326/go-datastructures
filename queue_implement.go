package main

import (
	"fmt"

	queue "github.com/kkumar326/go-dsa/pkgs/queue/linkedlist"
)

func queueRun() {
	var element queue.ItemType = "First element"
	queue := queue.NewQueue(&element)

	queue.Enqueue(1)
	queue.Enqueue("Second")
	queue.Enqueue(true)
	queue.Dequeue()
	queue.Enqueue(22)
	queue.Dequeue()
	//queue.Clear()

	fmt.Println(queue.Head())
}
