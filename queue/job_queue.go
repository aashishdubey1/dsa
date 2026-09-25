package queue

import (
	"fmt"
)

type Job struct {
	Id   int
	Desc string
}

type MessageQueue struct {
	queue    *ArrayQueue[Job]
	jobCount int
}

func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		queue: &ArrayQueue[Job]{
			items: []Job{},
		},
	}
}

func (mq *MessageQueue) Produce(desc string) {
	job := Job{Id: mq.jobCount + 1, Desc: desc}
	mq.queue.Enqueue(job)
	mq.jobCount++
}

func (mq *MessageQueue) Consume() {
	job, ok := mq.queue.Dequeue()
	if !ok {
		return
	}
	fmt.Println("job", job.Id, "consumed :", job.Desc)
}

func (mq *MessageQueue) ConsumeAll() {
	for !mq.queue.IsEmpty() {
		mq.Consume()
	}
}

func Run() {
	mq := NewMessageQueue()

	mq.Produce("send email")
	mq.Produce("send notificaiton")
	mq.Produce("send party")

	fmt.Println(mq.queue.items)
	mq.Consume()
	mq.Consume()
	mq.Consume()
	mq.Consume()
}
